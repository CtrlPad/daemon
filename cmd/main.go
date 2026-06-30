package main

import (
	"fmt"

	"github.com/ctrlpad/daemon/internal"
	"github.com/ctrlpad/daemon/internal/ble"
)

func main() {
	device, err := ble.ScanAndConnectToCtrlPad()
	if err != nil {
		fmt.Println(err)
		return
	}

	payload, err := ble.SetupNotifications(device)
	if err != nil {
		fmt.Println(err)
		return
	}

	for msg := range payload {
		internal.ExecuteAction(msg)
	}

	select {}
}
