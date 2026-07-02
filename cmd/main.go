package main

import (
	"github.com/ctrlpad/daemon/internal/ble"
	"github.com/ctrlpad/daemon/internal/executor"
	"github.com/ctrlpad/daemon/internal/logger"
)

func main() {
	device, err := ble.ScanAndConnectToCtrlPad()
	if err != nil {
		logger.Error("Connection", "err", err)
		return
	}

	payload, err := ble.SetupNotifications(device)
	if err != nil {
		logger.Error("SetupNotifications", "err", err)
		return
	}

	for msg := range payload {
		err := executor.ExecuteAction(msg)
		if err != nil {
			logger.Error("Executor", "err", err)
		}
	}

	select {}
}
