package internal

import (
	"encoding/json"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/ctrlpad/daemon/internal/logger"
)

type Action struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
	Exec string `json:"exec"`
	Icon string `json:"icon"`
}

func parseActionJson(actionString string) (Action, error) {
	var act Action
	err := json.Unmarshal([]byte(actionString), &act)
	if err != nil {
		return Action{}, err
	}

	return act, nil
}

func cleanDesktopCommand(cmdStr string) string {
	reg := regexp.MustCompile("%[uUfFkicm]")
	res := reg.ReplaceAllString(cmdStr, "")
	res = strings.TrimSpace(res)
	return res
}

func ExecuteAction(actionString string) error {
	action, err := parseActionJson(actionString)
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command(cleanDesktopCommand(action.Exec))
		logger.Info("Executor", "Executing", cleanDesktopCommand(action.Exec))
		err := cmd.Start()
		if err != nil {
			return err
		}

	default:
		logger.Error("Executor", "err", "OS not supported")
	}

	return nil
}
