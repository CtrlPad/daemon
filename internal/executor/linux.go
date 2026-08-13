package executor

import (
	"fmt"
)

func executeLinux(actionType string, target string) error {
	switch actionType {
	case "application":
		return execApplication(target)
	default:
		return fmt.Errorf("unknown action type: %s", actionType)
	}
}
