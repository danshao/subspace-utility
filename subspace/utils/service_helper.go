package utils

import "os/exec"

const (
	COMMAND_SERVICE = "service"
	COMMAND_SERVICE_START = "start"
	COMMAND_SERVICE_STOP = "stop"
)

// Start services by the given array sequence
func StartServices(services []string) error {
	size := len(services)
	for i := 0 ; i < size ; i++ {
		if _, err := exec.Command(COMMAND_SERVICE, services[i], COMMAND_SERVICE_START).Output(); nil != err {
			return err
		}
	}

	return nil
}

// Stop services by reversed sequence
func StopServices(services []string) error {
	size := len(services)
	for i := size - 1 ; i >= 0 ; i-- {
		if _, err := exec.Command(COMMAND_SERVICE, services[i], COMMAND_SERVICE_STOP).Output(); nil != err {
			return err
		}
	}

	return nil
}
