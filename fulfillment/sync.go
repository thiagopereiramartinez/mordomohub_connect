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
			DefaultNames: []string{"Canal 1"},
			Name:         "Canal 1",
		},
		WillReportState: false,
		RoomHint:        "Quarto",
		DeviceInfo:      manufacturer,
	}

	// Canal 2
	device_2 := device_1.Copy("2", "Canal 2")

	// Canal 3
	device_3 := device_1.Copy("3", "Canal 3")

	// Canal 4
	device_4 := device_1.Copy("4", "Canal 4")

	// Array com os dispositivos
	devices := []structs.Device{
		device_1,
		device_2,
		device_3,
		device_4,
	}

	c.JSON(fiber.Map{
		"requestId": requestId,
		"payload": fiber.Map{
			"agentUserId": 1,
			"devices":     devices,
		},
	})
}
