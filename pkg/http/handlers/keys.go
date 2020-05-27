package handlers

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gofiber/fiber"
)

// KeysHandler returns the server's public key in base64 format
func KeysHandler(c *fiber.Ctx) {
	key, err := ioutil.ReadFile("./keys/server_key_public_base64")

	if err != nil {
		fmt.Println("IO Error:", err.Error())

		c.Status(500).Send("Internal Server Error")
	} else {
		c.Send(strings.TrimSpace(string(key)))
	}
}
