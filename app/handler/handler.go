package handler

import (
	"github.com/gofiber/fiber/v2"
)

// @Tags        System
// @Summary     Ping Endpoint
// @Description Simple ping-pong endpoint to check if the server is running
// @Accept      json
// @Produce     json
// @Success     200 {object} object "Returns pong response"
// @Router      /ping [GET]
func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"ping": "pong"})
}

// @Tags        Welcome
// @Summary     Hello User
// @Description Endpoint to Welcome user and say Hello "Name"
// @Param       name query string true "Name in the URL param"
// @Accept      json
// @Produce     json
// @Success     200 {object} object "success"
// @Failure     400 {object} object "Request Error or parameter missing"
// @Failure     404 {object} object "When user not found"
// @Failure     500 {object} object "Server Error"
// @Router      /hello/:name [GET]
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello " + c.Params("name"))
}

// @Tags        System
// @Summary     Health Check
// @Description Endpoint to check the health status of the service
// @Accept      json
// @Produce     json
// @Success     200 {object} object "Returns service health status"
// @Failure     500 {object} object "Server Error"
// @Router      /health [GET]
func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// @Tags        Welcome
// @Summary     Index Page
// @Description Main entry point of the API
// @Accept      json
// @Produce     plain
// @Success     200 {string} string "Returns welcome message"
// @Router      / [GET]
func Index(c *fiber.Ctx) error {
	return c.SendString("Welcome To Soodeg API")
}
