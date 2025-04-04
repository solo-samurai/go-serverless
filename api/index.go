package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/solo-samurai/go-serverless/app/routes"
)

var (
	app *fiber.App
)

// @title Golang Vercel Serverless API
// @version 1.0
// @description This is a Golang Vercel Serverless API.
// @description This API is built using the Fiber framework and is designed to be deployed on Vercel.

// @schemes https http
func init() {
	app = fiber.New()
	routes.Register(app)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app).ServeHTTP(w, r)
}
