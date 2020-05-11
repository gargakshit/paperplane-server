package http

import (
	"log"
	"sync"

	"github.com/gargakshit/paperplane-server/pkg/http/handlers"
	"github.com/gargakshit/paperplane-server/pkg/http/handlers/directory"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
)

// BootstrapHTTP bootstraps the HTTP Server
func BootstrapHTTP(listenAddress string, wg *sync.WaitGroup) {
	defer wg.Done()

	httpServer := fiber.New(&fiber.Settings{
		ServerHeader: "paperplane-v2",
	})

	httpServer.Use(recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString("Internal Server Error")
			c.SendStatus(500)
		},
	}))

	httpServer.Get("/config", handlers.ClusterConfigHandler)

	directoryGroup := httpServer.Group("/directory")
	directoryGroup.Post("/register", directory.RegisterHandler)

	log.Fatalln(httpServer.Listen(listenAddress))
}
