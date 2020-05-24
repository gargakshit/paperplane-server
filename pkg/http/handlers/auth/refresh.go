package auth

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

// RefreshToken refreshes the auth token for the client
func RefreshToken(c *fiber.Ctx) {
	// authHeader := string(c.Fasthttp.Request.Header.Peek("Authorization")[:])
	refreshToken := string(c.Fasthttp.Request.Header.Peek("X-Refresh-Token")[:])

	token, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err == nil && token.Valid {
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
				Audience: claims.Audience,
			})

			newTokenString, err := newToken.SignedString([]byte(os.Getenv("JWT_KEY")))

			if err != nil {
				log.Println("JWT Error:", err.Error())

				c.Status(500).Send("Internal Server Error")

				return
			}

			c.Send(newTokenString)
		} else {
			log.Println("JWT Error:", err.Error())

			c.Status(500).Send("Internal Server Error")

			return
		}
	} else {
		c.Status(401).Send("Unauthorized")
	}
}
