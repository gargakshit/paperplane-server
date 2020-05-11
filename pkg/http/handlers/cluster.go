package handlers

import (
	"github.com/gargakshit/paperplane-server/config"
	"github.com/gofiber/fiber"
)

// ClusterConfigHandler sends the server config
func ClusterConfigHandler(ctx *fiber.Ctx) {
	ctx.JSON(config.GlobalConfig.ClusterConfig)
}
