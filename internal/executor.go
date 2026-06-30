package internal

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
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
	fmt.Println(action.Exec)

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command(cleanDesktopCommand(action.Exec))
		fmt.Println("Executing: ", cleanDesktopCommand(action.Exec))
		err := cmd.Start()
		if err != nil {
			return err
		}

	default:
		fmt.Println("Operating system not supported :(")
	}

	return nil
}
