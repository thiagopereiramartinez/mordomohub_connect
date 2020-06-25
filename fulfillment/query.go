package fulfillment

import (
	"encoding/json"
	"github.com/gofiber/fiber"
)

func Query(c *fiber.Ctx, requestId string, payload map[string]interface{}) error {

	str, _ := json.Marshal(payload["devices"])
	devices := make([]map[string]interface{}, 0, 10)

	json.Unmarshal(str, &devices)

	for _, device := range devices {

		switch device["id"] {
		// Termostato
		case "5":
			c.JSON(fiber.Map{
				"requestId": requestId,
				"payload": map[string] interface{} {
					"devices": map[string] interface{} {
						"5": map[string] interface{} {
							"online": true,
							"activeThermostatMode": "none",
							"thermostatTemperatureAmbient": 25.3,
							"thermostatHumidityAmbient": 45.3,
							"status": "SUCCESS",
						},
					},
				},
			})
		}
	}

	return nil
}
