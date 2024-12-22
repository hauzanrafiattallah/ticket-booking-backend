package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hauzanrafiattallah/ticket-booking-backend/handlers"
	"github.com/hauzanrafiattallah/ticket-booking-backend/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking",
		ServerHeader: "Fiber",
	})

	// Repositories
	EventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/events"), EventRepository)

	app.Listen(":3000")
}
