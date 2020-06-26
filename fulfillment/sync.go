package fulfillment

import (
	"github.com/gofiber/fiber"
	"github.com/thiagopereiramartinez/mordomohub_connect/structs"
)

func Sync(c *fiber.Ctx, requestId string, payload map[string]interface{}) {

	// Fabricante
	manufacturer := structs.DeviceInfo{
		Manufacturer: "Thiago P. Martinez",
		Model:        "MordomoHub",
		HwVersion:    "1.0.0",
		SwVersion:    "1.0.0",
	}

	// Canal 1
	device_1 := structs.Device{
		Id:     "1",
		Type:   structs.TYPE_OUTLET,
		Traits: []string{structs.TRAITS_ON_OFF},
		Name: structs.DeviceName{
			DefaultNames: []string{"Tomada 1"},
			Name:         "Tomada 1",
		},
		WillReportState: false,
		RoomHint:        "Quarto",
		DeviceInfo:      manufacturer,
		Attributes: map[string]interface{}{
			"commandOnlyOnOff": true,
		},
	}

	// Canal 2
	device_2 := device_1.Copy("2", "Tomada 2")

	// Canal 3
	device_3 := device_1.Copy("3", "Tomada 3")

	// Canal 4
	device_4 := device_1.Copy("4", "Tomada 4")

	// Termostato
	device_5 := structs.Device{
		Id:   "5",
		Type: structs.TYPE_SENSOR,
		Traits: []string{
			structs.TRAITS_HUMIDITY_SETTING,
			structs.TRAITS_TEMPERATURE_CONTROL,
		},
		Name: structs.DeviceName{
			DefaultNames: []string{"DHT22"},
			Name:         "DHT22",
		},
		WillReportState: true,
		RoomHint:        "Quarto",
		DeviceInfo:      manufacturer,
		Attributes: map[string]interface{}{
			"queryOnlyHumiditySetting":    true,
			"queryOnlyTemperatureControl": true,
			"temperatureUnitForUX":        "C",
		},
	}

	// Array com os dispositivos
	devices := []structs.Device{
		device_1,
		device_2,
		device_3,
		device_4,
		device_5,
	}

	c.JSON(fiber.Map{
		"requestId": requestId,
		"payload": fiber.Map{
			"agentUserId": "1",
			"devices":     devices,
		},
	})
}
