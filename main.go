package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/thiagopereiramartinez/mordomo_connect/auth"
	"github.com/thiagopereiramartinez/mordomo_connect/fulfillment"
	"github.com/thiagopereiramartinez/mordomo_connect/sensors"
	"os"
	"strconv"
)

func main() {

	// Criar aplicação
	app := fiber.New()

	// Configurar middlewares
	app.Use(helmet.New())

	// Autenticação
	app.Get("/auth", auth.Auth)
	app.Post("/token", auth.Token)

	// Fulfillment
	app.Post("/fulfillment", fulfillment.Fulfillment)

	// Sensors
	app.Get("/sensors", sensors.Sensors)

	// Iniciar serviço
	port := 8080
	if env_port := os.Getenv("PORT"); env_port != "" {
		if p, err := strconv.Atoi(env_port); err != nil {
			fmt.Errorf("%v", err)
			os.Exit(1)
		} else {
			port = p
		}
	}

	app.Listen(port)
}
