package utils

import "os/exec"

// check if linux binary exists in path
func CheckLinuxBinary(binary string) error {
	_, err := exec.LookPath(binary)
	return err
}
