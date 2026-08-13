package executor

import "os/exec"

func execApplication(target string) error {
	cmd := exec.Command(target)
	err := cmd.Start()
	if err != nil {
		return err
	}
	go cmd.Wait()
	return nil
}
