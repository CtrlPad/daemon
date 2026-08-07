package executor

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type Button struct {
	ID     int16  `json:"id"`
	Name   string `json:"name"`
	Action string `json:"action"`
	Icon   string `json:"icon"`
}

func parseButtonConfig(buttonConfig string) (Button, error) {
	var btn Button
	err := json.Unmarshal([]byte(buttonConfig), &btn)
	if err != nil {
		return Button{}, err
	}
	return btn, nil
}

// Example: application:firefox
func parseActionPrefix(action string) (actionType string, target string, err error) {
	parts := strings.Split(action, ":")
	actionType = parts[0]
	target = parts[1]

	if actionType == "" || target == "" {
		return "", "", fmt.Errorf("Error parsing action")
	}
	return actionType, target, nil
}

func ExecuteAction(buttonConfig string) error {
	btnConfig, err := parseButtonConfig(buttonConfig)
	if err != nil {
		return err
	}
	actionType, target, err := parseActionPrefix(btnConfig.Action)
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		return executeLinux(actionType, target)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}
