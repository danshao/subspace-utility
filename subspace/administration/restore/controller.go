package restore

import (
	"sync"
	"errors"
	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/vpn"
)

var instance *controller
var once sync.Once

var mutex sync.Mutex

type Callback interface {
	OnStart()
	OnCancel()
	OnSuccess(string)
	OnFail(error)
}

type Step int

const (
	IDLE      = iota
	RUNNING
	SUCCEED
	FAILED
	CANCELING
	CANCELED
	UNKNOWN
)

type Status struct {
	Step   Step
	Result string
	Error  error
}

// Define service start sequence. Service monit should be started at the end.
var SERVICES = []string{
	"vpnsession",
	"vpnprofile",
	"vpnserver",
	"monit",
}

func GetInstance() *controller {
	once.Do(func() {
		instance = &controller{}
		instance.init()
	})
	return instance
}

type controller struct {
	running  bool
	canceled bool

	ignoreCheckSum bool

	dbUri string
	db    *gorm.DB

	callback Callback
	result   string
	err      error
}

func (controller *controller) init() {
	controller.dbUri = config.GetDefaultMysqlUri()
}

func (controller *controller) SetIgnoreCheckSum(ignore bool) {
	controller.ignoreCheckSum = ignore
}

func (controller *controller) IsIgnoreCheckSum() bool {
	return controller.ignoreCheckSum
}

func (controller *controller) SetDatabaseUri(uri string) {
	controller.dbUri = uri
}

func (controller *controller) SetCallback(callback Callback) {
	controller.callback = callback
}

func (controller *controller) RemoveCallback(callback Callback) {
	controller.callback = nil
}

func (controller *controller) Start(path string) (success bool) {
	mutex.Lock()
	defer mutex.Unlock()

	if controller.running {
		return false
	}

	//Reset all status before run
	controller.canceled = false
	controller.result = ""
	controller.err = nil

	controller.running = true
	controller.onStart()

	go controller.run(path)
	return true
}

func (controller *controller) GetStatus() Status {
	status := Status{}
	switch {
	case !controller.running && nil == controller.err && "" == controller.result && !controller.canceled:
		status.Step = IDLE

	case controller.canceled && controller.running:
		status.Step = CANCELING

	case controller.canceled && !controller.running:
		status.Step = CANCELED

	case !controller.running && nil != controller.err:
		status.Step = FAILED
		status.Error = controller.err

	case !controller.running && "" != controller.result:
		status.Step = SUCCEED
		status.Result = controller.result

	case controller.running:
		status.Step = RUNNING

	default: // Should never happen
		status.Step = UNKNOWN
	}
	return status
}

func (controller *controller) IsRunning() bool {
	return controller.running
}

func (controller *controller) Cancel() {
	controller.canceled = true
}

func (controller *controller) IsCanceled() bool {
	return controller.canceled
}

func (controller *controller) run(path string) {
	defer controller.cleanup()

	if controller.canceled {
		controller.onCancel()
		return
	}

	// Read from path
	yamlData, err := utils.ReadFromFile(path)
	if nil != err {
		controller.onFail(err)
		return
	}

	cfg, err := administration.ParseConfig(yamlData)
	if nil != err {
		controller.onFail(err)
		return
	}

	// Test checksum if needed
	if !controller.ignoreCheckSum && !cfg.IsCheckSumMatch() {
		controller.onFail(errors.New("Checksum mismatch, or set controller.SetIgnoreCheckSum to true"))
		return
	}

	// Validate cfg
	db, err := gorm.Open("mysql", controller.dbUri)
	defer db.Close()
	if err := cfg.Validate(db); nil != err {
		controller.onFail(err)
		return
	}

	// Start Restore

	// Stop monit, session daemon, profile daemon
	utils.StopServices(SERVICES)
	// Restart services if any error
	defer utils.StartServices(SERVICES)

	// Restore DB data
	if nil != err {
		controller.onFail(err)
		return
	}

	// Lock write operation on necessary tables.
	if err := utils.LockTableWrite(db,
		model.User{}.TableName(),
		model.Profile{}.TableName(),
		model.System{}.TableName(),
		model.Log{}.TableName(),
		model.ProfileSnapshot{}.TableName(),
	); nil != err {
		controller.onFail(err)
		return
	}
	// Unlock tables after all
	defer utils.UnlockTable(db)

	// Truncate all table
	utils.TruncateTable(db,
		model.User{}.TableName(),
		model.Profile{}.TableName(),
		// Usage data do NOT keep.
		model.ProfileSnapshot{}.TableName(),
		model.Log{}.TableName(),
	)

	//TODO Check profile related user_id is exist or not.

	// Update system data
	sys := cfg.GetSystem()
	db.Table(sys.TableName()).Updates(sys.DataToRestore())

	// Insert user data
	for _, user := range cfg.GetUsers() {
		db.Table(user.TableName()).Create(&user)
	}

	// Insert profile data and convert to vpn account
	profiles := cfg.GetProfiles()
	accounts := make([]vpn.Account, 0)
	for _, profile := range profiles {
		accounts = append(accounts, ToAccount(profile))
		db.Table(profile.TableName()).Create(&profile)
	}

	// Format softether cfg
	vpnServer := vpn.Softether{
		AdministrationPort: config.DEFAULT_VPN_SERVER_ADMINISTRATION_PORT,
		AdminPassword: config.DEFAULT_VPN_SERVER_ADMINISTRATION_PASSWORD,
		PreSharedKey: sys.PreSharedKey,
		Hub: vpn.Hub{
			Name: config.DEFAULT_HUB_NAME,
			Accounts: accounts,
		},
	}
	softetherConfig, err := vpn.GenerateSoftetherConfig(vpnServer)
	if nil != err {
		controller.onFail(err)
		return
	}

	// Write Softether config
	if err := utils.WriteToFile(config.SOFTETHER_CONFIG_PATH, softetherConfig); nil != err {
		controller.onFail(err)
		return
	}

	//Finished do all defer call than call onSuccess
	defer controller.onSuccess(path)
}

func (controller *controller) cleanup() {
	controller.running = false
}

func (controller *controller) onStart() {
	if nil != controller.callback {
		controller.callback.OnStart()
	}
}

func (controller *controller) onCancel() {
	if nil != controller.callback {
		controller.callback.OnCancel()
	}
}

func (controller *controller) onSuccess(result string) {
	controller.result = result
	if nil != controller.callback {
		controller.callback.OnSuccess(result)
	}
}

func (controller *controller) onFail(e error) {
	controller.err = e
	if nil != controller.callback {
		controller.callback.OnFail(e)
	}
}

func ToAccount(p model.Profile) vpn.Account {
	profile := vpn.Account{
		Username:       p.UserName,
		PasswordHash:   p.PasswordHash,
		NtLmSecureHash: p.NtLmSecureHash,
		RawRealName:    p.FullName,
		RawNote:        p.Description,
		LoginCount:     p.LoginCount,
		RevokedTime:    p.RevokedDate,
		LastLoginTime:  p.LastLoginDate,
		UpdatedTime:    p.UpdatedDate,
		CreatedTime:    p.CreatedDate,
	}
	return profile
}
