package action

import (
	"os/exec"

	"github.com/ctrlpad/daemon/internal/utils"
)

func ExecApplication(target string) error {
	err := utils.CheckLinuxBinary(target)
	if err != nil {
		return err
	}
	cmd := exec.Command(target)
	err = cmd.Start()
	if err != nil {
		return err
	}
	go cmd.Wait()
	return nil
}
