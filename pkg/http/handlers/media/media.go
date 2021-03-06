package media

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gargakshit/paperplane-server/utils"
	"github.com/gofiber/fiber"
)

// UploadFile is the handler to upload files to the server
func UploadFile(c *fiber.Ctx) {
	authHeader := string(c.Fasthttp.Request.Header.Peek("Authorization")[:])

	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err == nil && token.Valid {
		file, err := c.FormFile("file")

		if err == nil {
			id := utils.GenerateFileID()
			c.SaveFile(file, fmt.Sprintf("./uploads/%s", id))
			c.Send(id)
		} else {
			c.Status(500).Send("Internal Server Error")
		}
	} else {
		c.Status(401).Send("Unauthorized")
	}
}
