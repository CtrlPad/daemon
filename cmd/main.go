package main

import (
	"github.com/charmbracelet/log"
	"github.com/ctrlpad/daemon/internal/ble"
	"github.com/ctrlpad/daemon/internal/executor"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	log.Info("App Info", "Version", version, "Commit", commit, "Date", date)
	device, err := ble.ScanAndConnectToCtrlPad()
	if err != nil {
		log.Error("Connection", "err", err)
		return
	}

	payload, err := ble.SetupNotifications(device)
	if err != nil {
		log.Error("SetupNotifications", "err", err)
		return
	}

	for msg := range payload {
		err := executor.ExecuteAction(msg)
		if err != nil {
			log.Error("Executor", "err", err)
		}
	}

	select {}
}
