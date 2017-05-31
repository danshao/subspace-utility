package main

import (
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration/restore"
	"time"
	"testing"
)

const RESTORE_PATH = "/tmp/subspace.config"

type MyRestoreCallback struct {}

func (c MyRestoreCallback) OnStart() {
	fmt.Println("restore OnStart")
}

func (c MyRestoreCallback) OnCancel() {
	fmt.Println("restore OnCancel")
}

func (c MyRestoreCallback) OnSuccess(yaml string) {
	fmt.Println("restore OnSuccess to", RESTORE_PATH)
}

func (c MyRestoreCallback) OnFail(e error) {
	fmt.Println("restore callback OnFail.")
}

func TestRestore(*testing.T) {
	restoreController := restore.GetInstance()
	restoreController.SetCallback(MyRestoreCallback{})  // Not necessary
	restoreController.Start(RESTORE_PATH)
	
	ticker := time.NewTicker(time.Millisecond * 200)
	for t := range ticker.C {
		status := restoreController.GetStatus()
		switch status.Step {
		case restore.IDLE:
			fmt.Println(t, "restore idle")
		case restore.RUNNING:
			fmt.Println(t, "restore running")
		case restore.SUCCEED:
			fmt.Println(t, "restore succeed", status.Result)
		case restore.FAILED:
			fmt.Println(t, "restore failed", status.Error)
		case restore.CANCELING:
			fmt.Println(t, "restore canceling")
		case restore.CANCELED:
			fmt.Println(t, "restore canceled")
		case restore.UNKNOWN:
			fmt.Println(t, "something wrong")
		}

		isRunning := restoreController.IsRunning()
		if isRunning {
		} else {
			ticker.Stop()
			break
		}
	}
}