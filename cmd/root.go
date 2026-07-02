package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/ctrlpad/daemon/internal/ble"
	"github.com/ctrlpad/daemon/internal/executor"
)

func Run() int {
	device, err := ble.ScanAndConnectToCtrlPad()
	if err != nil {
		log.Error("Connection", "err", err)
		return 1
	}

	payload, err := ble.SetupNotifications(device)
	if err != nil {
		log.Error("SetupNotifications", "err", err)
		return 1
	}

	for msg := range payload {
		err := executor.ExecuteAction(msg)
		if err != nil {
			log.Error("Executor", "err", err)
		}
	}
	return 1
}
