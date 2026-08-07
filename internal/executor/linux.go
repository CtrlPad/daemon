package executor

import (
	"fmt"
	"os/exec"
)

func executeLinux(actionType string, target string) error {
	switch actionType {
	case "application":
		cmd := exec.Command(target)
		err := cmd.Start()
		if err != nil {
			return err
		}
		go cmd.Wait()
	default:
		return fmt.Errorf("unknown action type: %s", actionType)
	}
	return nil
}
