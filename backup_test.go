package main

import (
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration/backup"
	"time"
	"testing"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils"
)

const BACKUP_PATH = "/tmp/subspace.config"

type MyBackupCallback struct {}

func (c MyBackupCallback) OnStart() {
	fmt.Println("backup OnStart")
}

func (c MyBackupCallback) OnCancel() {
	fmt.Println("backup OnCancel")
}

func (c MyBackupCallback) OnSuccess(yaml string) {
	fmt.Println("backup OnSuccess")
	if e := utils.WriteToFile(BACKUP_PATH, yaml); nil != e {
		fmt.Println("Write to", BACKUP_PATH, "fail.")
	} else {
		fmt.Println("Write to", BACKUP_PATH, "success.")
	}
}

func (c MyBackupCallback) OnFail(e error) {
	fmt.Println("backup callback OnFail.")
}

func TestBackup(*testing.T) {
	backupController := backup.GetInstance()
	backupController.SetCallback(MyBackupCallback{})  // Not necessary
	backupController.Start()
	
	ticker := time.NewTicker(time.Millisecond * 200)
	for t := range ticker.C {
		status := backupController.GetStatus()
		switch status.Step {
		case backup.IDLE:
			fmt.Println(t, "backup idle")
		case backup.RUNNING:
			fmt.Println(t, "backup running")
		case backup.SUCCEED:
			fmt.Println(t, "backup succeed", status.Result)
		case backup.FAILED:
			fmt.Println(t, "backup failed", status.Error)
		case backup.CANCELING:
			fmt.Println(t, "backup canceling")
		case backup.CANCELED:
			fmt.Println(t, "backup canceled")
		case backup.UNKNOWN:
			fmt.Println(t, "something wrong")
		}

		isRunning := backupController.IsRunning()
		if isRunning {
		} else {
			ticker.Stop()
			break
		}
	}
}