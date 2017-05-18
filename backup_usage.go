package main

import (
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration/backup"
	"time"
)

const BACKUP_PATH = "/tmp/subspace.config"

type MyBackupCallback struct {}

func (c MyBackupCallback) OnStart() {
	fmt.Println("backup start")
}

func (c MyBackupCallback) OnCancel() {
	fmt.Println("backup cancel")
}

func (c MyBackupCallback) OnSuccess(yaml string) {
	fmt.Println("backup success to", BACKUP_PATH)
	fmt.Println(yaml)
}

func (c MyBackupCallback) OnFail(e error) {
	fmt.Println("backup fail.", e)
}

func main() {
	backupController := backup.GetInstance()
	backupController.Start(BACKUP_PATH, MyBackupCallback{})
	ticker := time.NewTicker(time.Millisecond * 200)
	for t := range ticker.C {
		isRunning := backupController.IsBackingUp()
		if isRunning {
			fmt.Println("backup running", t)
		} else {
			fmt.Println("backup stop", t)
			ticker.Stop()
		}
	}
}