package fulfillment

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber"
	"log"
	"time"
)

func Query(c *fiber.Ctx, requestId string, payload map[string]interface{}) error {

	str, _ := json.Marshal(payload["devices"])
	devices := make([]map[string]interface{}, 0, 10)

	json.Unmarshal(str, &devices)

	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "masterdeveloper-mordomo-hub"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	m := make(map[string]interface{})

	for _, device := range devices {

		switch device["id"] {
		// Sensor temperatura/umidade
		case "5":

			// Ler do Firestore
			snap, _ := client.Collection("esp32").Doc("state").Get(ctx)
			doc := snap.Data()
			temp, humidity, _ := doc["temp"], int(doc["humidity"].(float64)), doc["timestamp"].(time.Time)

			m["5"] = map[string]interface{}{
				"online":                     true,
				"humidityAmbientPercent":     humidity,
				"temperatureSetpointCelsius": temp,
				"temperatureAmbientCelsius":  temp,
				"status":                     "SUCCESS",
			}
		}
	}

	c.JSON(fiber.Map{
		"requestId": requestId,
		"payload": map[string]interface{}{
			"devices": m,
		},
	})

	return nil
}
