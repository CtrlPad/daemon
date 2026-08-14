package executor

import (
	"fmt"

	"github.com/ctrlpad/daemon/internal/action"
)

func executeLinux(actionType string, target string) error {
	switch actionType {
	case "application":
		return action.ExecApplication(target)
	case "volume":
		return action.ExecVolume(target)
	default:
		return fmt.Errorf("unknown action type: %s", actionType)
	}
}
