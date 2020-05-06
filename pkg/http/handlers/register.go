package handlers

import (
	"github.com/gofiber/fiber"
)

// RegisterHandler is the HTTP handler for the /register endpoint
func RegisterHandler(ctx *fiber.Ctx) {
	ctx.Send("OK")
}
