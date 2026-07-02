package ble

import (
	"github.com/charmbracelet/log"

	"tinygo.org/x/bluetooth"
)

func ScanAndConnectToCtrlPad() (*bluetooth.Device, error) {
	err := Adapter.Enable()
	if err != nil {
		log.Error("Adapter", "err", err)
	}
	deviceChan := make(chan bluetooth.ScanResult, 1)

	log.Info("Scanning...")
	err = Adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		log.Info("Found device", "Device Name", result.LocalName(), "RSSI", result.RSSI, "Address", result.Address.String())
		if result.LocalName() == "ctrlPad_BLE" {
			adapter.StopScan()
			deviceChan <- result
		}
	})
	if err != nil {
		return nil, err
	}

	foundDevice := <-deviceChan

	device, err := Adapter.Connect(foundDevice.Address, bluetooth.ConnectionParams{})
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func SetupNotifications(device *bluetooth.Device) (chan string, error) {
	srvcs, err := device.DiscoverServices([]bluetooth.UUID{CtrlPadServiceUUID})
	if err != nil {
		return nil, err
	}
	srvc := srvcs[0]

	chars, err := srvc.DiscoverCharacteristics([]bluetooth.UUID{CtrlPadCharacteristicUUID})
	if err != nil {
		return nil, err
	}
	char := chars[0]
	log.Info("Found characteristic", "UUID", char.UUID().String())

	notifyChan := make(chan string, 1)

	err = char.EnableNotifications(func(buf []byte) {
		notifyChan <- string(buf)
	})
	if err != nil {
		return nil, err
	}

	return notifyChan, nil
}
