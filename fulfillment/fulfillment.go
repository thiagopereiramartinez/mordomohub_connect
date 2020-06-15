package fulfillment

import "github.com/gofiber/fiber"

type FulfillmentRequest struct {
	RequestId string              `json:"requestId"`
	Inputs    []FulfillmentIntent `json:"inputs"`
}

type FulfillmentIntent struct {
	Intent  string                 `json:"intent"`
	Payload map[string]interface{} `json:"payload"`
}

const (
	SYNC    = "action.devices.SYNC"
	EXECUTE = "action.devices.EXECUTE"
)

func Fulfillment(c *fiber.Ctx) {
	request := new(FulfillmentRequest)

	if err := c.BodyParser(request); err != nil {
		c.Status(500).Send("Internal server error")
	}

	for _, input := range request.Inputs {
		switch input.Intent {
		case SYNC:
			Sync(c, request.RequestId, input.Payload)
		case EXECUTE:
			Execute(c, request.RequestId, input.Payload)
		}
	}

}
