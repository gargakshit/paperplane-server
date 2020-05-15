package http

import (
	"log"
	"sync"

	"github.com/gargakshit/paperplane-server/pkg/http/handlers"
	"github.com/gargakshit/paperplane-server/pkg/http/handlers/directory"
	"github.com/gargakshit/paperplane-server/pkg/http/handlers/media"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
)

// BootstrapHTTP bootstraps the HTTP Server
func BootstrapHTTP(listenAddress string, wg *sync.WaitGroup) {
	defer wg.Done()

	httpServer := fiber.New(&fiber.Settings{
		ServerHeader: "paperplane-v2",
		BodyLimit:    1024 * 1024 * 65, // 65 mb
	})

	httpServer.Use(recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			log.Println("HTTP Error:", err)
			c.SendString("Internal Server Error")
			c.SendStatus(500)
		},
	}))

	httpServer.Get("/config", handlers.ClusterConfigHandler)

	directoryGroup := httpServer.Group("/directory")
	directoryGroup.Post("/register", directory.RegisterHandler)
	directoryGroup.Get("/:id", directory.GetUserPublicKey)

	mediaGroup := httpServer.Group("/media")
	mediaGroup.Post("/upload", media.UploadFile)
	mediaGroup.Static("/get", "./uploads")

	log.Fatalln(httpServer.Listen(listenAddress))
}
