package action

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/ctrlpad/daemon/internal/utils"
)

func parseVolumeTarget(target string) (string, error) { // +10 => 10%+
	value := target
	suffix := "%"
	switch target[0] {
	case '+':
		value = target[1:]
		suffix = "%+"
	case '-':
		value = target[1:]
		suffix = "%-"
	}

	amount, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", fmt.Errorf("invalid volume target: %s", target)
	}
	if amount < 0 {
		return "", fmt.Errorf("invalid volume target: %s", target)
	}

	return value + suffix, nil
}

func ExecVolume(target string) error {
	err := utils.CheckLinuxBinary("wpctl")
	if err != nil {
		return err
	}
	arg, err := parseVolumeTarget(target)
	if err != nil {
		return err
	}
	cmd := exec.Command("wpctl", "set-volume", "@DEFAULT_AUDIO_SINK@", arg)
	err = cmd.Start()
	if err != nil {
		return err
	}
	go cmd.Wait()
	return nil
}
