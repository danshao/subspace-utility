package backup

import (
	"sync"
	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration"
	"os"
	"bufio"
	"errors"
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

func GetInstance() *controller {
	once.Do(func() {
		instance = &controller{}
		instance.init()
	})
	return instance
}

type controller struct {
	running      bool
	canceled     bool

	dbUri string
	db    *gorm.DB
}

func (controller *controller) init() {
	controller.dbUri = config.GetDefaultMysqlUri()
}

func (controller *controller) SetDatabaseUri(uri string) {
	controller.dbUri = uri
}

func (controller *controller) Start(path string, callback Callback) {
	mutex.Lock()
	defer mutex.Unlock()

	if controller.running {
		go callback.OnFail(errors.New("Back up process is currently running."))
		return
	}

	controller.canceled = false
	controller.running = true

	go callback.OnStart()

	go controller.run(path, callback)
}

func (controller *controller) IsBackingUp() bool {
	return controller.running
}

func (controller *controller) Cancel() {
	controller.canceled = true
}

func (controller *controller) run(path string, callback Callback) {
	defer controller.cleanup()

	if controller.canceled {
		go callback.OnCancel()
		return
	}

	data, err := administration.GenerateConfig(controller.dbUri)
	if nil != err {
		go callback.OnFail(err)
		return
	}

	if controller.canceled {
		go callback.OnCancel()
		return
	}

	if e := writeToFile(path, data); nil != e {
		go callback.OnFail(e)
		return
	}

	go callback.OnSuccess(data)
}

func writeToFile(path string, data string) error {
	file, err := os.Create(path)
	defer file.Close()

	if nil != err {
		return err
	}

	bufferWriter := bufio.NewWriter(file)
	if _, e := bufferWriter.WriteString(data); nil != e {
		return e
	}

	if e := bufferWriter.Flush(); nil != e {
		return e
	}

	return nil
}

func (controller *controller) cleanup() {
	controller.running = false
}
