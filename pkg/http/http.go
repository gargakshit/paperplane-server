package http

import (
	"log"
	"sync"

	"github.com/gofiber/fiber"
)

// BootstrapHTTP bootstraps the HTTP Server
func BootstrapHTTP(listenAddress string, wg *sync.WaitGroup) {
	defer wg.Done()

	httpServer := fiber.New(&fiber.Settings{
		ServerHeader: "paperplane-v2",
	})

	httpServer.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello")
	})

	log.Fatalln(httpServer.Listen(listenAddress))
}
