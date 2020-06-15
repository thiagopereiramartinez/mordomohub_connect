package fulfillment

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/thiagopereiramartinez/mordomohub_connect/structs"
	cloudiot "google.golang.org/api/cloudiot/v1"
)

const (
	GCLOUD_PROJECT_ID   string = "masterdeveloper-mordomo-hub"
	GCLOUD_IOT_REGION   string = "us-central1"
	GCLOUD_IOT_REGISTRY string = "mordomo"
	GCLOUD_IOT_DEVICE   string = "esp32"
)

func Execute(c *fiber.Ctx, requestId string, payload map[string]interface{}) error {

	str, _ := json.Marshal(payload["commands"])
	commands := new([]structs.Command)

	json.Unmarshal(str, &commands)

	ctx := context.Background()
	client, err := cloudiot.NewService(ctx)
	if err != nil {
		return nil
	}

	name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", GCLOUD_PROJECT_ID, GCLOUD_IOT_REGION, GCLOUD_IOT_REGISTRY, GCLOUD_IOT_DEVICE)

	resultCommands := make([]map[string]interface{}, 0, 10)

	for _, c := range *commands {
		for _, e := range c.Execution {
			for _, d := range c.Devices {

				on := e.Params["on"].(bool)
				cmd := ""
				if on {
					cmd = "ON"
				} else {
					cmd = "OFF"
				}
				req := cloudiot.SendCommandToDeviceRequest{
					BinaryData: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s-%s", cmd, d.Id))),
				}

				_, err := client.Projects.Locations.Registries.Devices.SendCommandToDevice(name, &req).Do()
				if err != nil {
					resultCommands = append(resultCommands, map[string] interface{} {
						"ids": []string { d.Id },
						"status": "ERROR",
						"errorCode": "deviceTurnedOff",
					})
					continue
				}

				resultCommands = append(resultCommands, map[string] interface{} {
					"ids": []string { d.Id },
					"status": "SUCCESS",
					"states": map[string] interface{} {
						"on": on,
						"online": true,
					},
				})
			}
		}
	}

	c.JSON(fiber.Map{
		"requestId": requestId,
		"payload": map[string] interface{} {
			"commands": resultCommands,
		},
	})

	return nil
}
