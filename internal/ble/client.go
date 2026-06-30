package ble

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

func ScanAndConnectToCtrlPad() (*bluetooth.Device, error) {
	err := Adapter.Enable()
	if err != nil {
		fmt.Println("Error enabeling Adapter: ", err)
	}
	deviceChan := make(chan bluetooth.ScanResult, 1)

	println("scanning...")
	err = Adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		println("Found device:", result.Address.String(), result.RSSI, result.LocalName())
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
		fmt.Println("Error connecting to device:", err)
		return nil, err
	}
	return &device, nil
}

func SetupNotifications(device *bluetooth.Device) (chan string, error) {
	srvcs, err := device.DiscoverServices([]bluetooth.UUID{CtrlPadServiceUUID})
	if err != nil {
		fmt.Println(err)
	}
	if len(srvcs) == 0 {
		panic("could not find heart rate service")
	}
	srvc := srvcs[0]

	chars, err := srvc.DiscoverCharacteristics([]bluetooth.UUID{CtrlPadCharacteristicUUID})
	if err != nil {
		println(err)
	}
	if len(chars) == 0 {
		panic("could not find heart rate characteristic")
	}
	char := chars[0]
	println("found characteristic", char.UUID().String())

	notifyChan := make(chan string, 1)

	err = char.EnableNotifications(func(buf []byte) {
		notifyChan <- string(buf)
	})
	if err != nil {
		return nil, err
	}

	return notifyChan, nil
}
