package backup

import (
	"sync"

	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
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
	IDLE = iota
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

	dbUri string
	db    *gorm.DB

	callback Callback
	result   string
	err      error
}

func (controller *controller) init() {
	controller.dbUri = config.GetDefaultMysqlUri()
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

func (controller *controller) Start() (success bool) {
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

	go controller.run()
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

func (controller *controller) run() {
	defer controller.cleanup()

	if controller.canceled {
		controller.onCancel()
		return
	}

	data, err := administration.GenerateConfig(controller.dbUri)
	if nil != err {
		controller.onFail(err)
		return
	}

	if controller.canceled {
		controller.onCancel()
		return
	}

	defer controller.onSuccess(data)
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
