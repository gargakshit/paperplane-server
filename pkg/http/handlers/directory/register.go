package directory

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gargakshit/paperplane-server/database"
	"github.com/gargakshit/paperplane-server/model"
	"github.com/gargakshit/paperplane-server/utils"
	"github.com/gofiber/fiber"
)

// RegisterHandler is the HTTP handler for the /register endpoint
func RegisterHandler(ctx *fiber.Ctx) {
	var body model.RegistrationBody
	err := ctx.BodyParser(&body)

	if err != nil {
		log.Panicln("HTTP Error:", err.Error())

		ctx.Status(500).Send("Internal Server Error")
	} else if body.PubKey == "" {
		ctx.Status(400).Send("Bad Request")
	} else {
		if err != nil {
			log.Println("HTTP Error:", err.Error())

			ctx.Status(500).Send("Internal Server Error")
		} else {
			if utils.IsBase64Valid(body.PubKey) {
				id := utils.GenerateRandomID()

				if err != nil {
					log.Println("DB Error:", err.Error())

					ctx.Status(500).Send("Internal Server Error")
				}

				directoryCollection := database.MongoConnection.Database("paperplane").Collection("directory")

				var resultUser model.UserDataType

				if err = directoryCollection.FindOne(context.TODO(), &model.UserDataType{
					PubKey: body.PubKey,
				}).Decode(&resultUser); err != nil {
					_, err = directoryCollection.InsertOne(context.TODO(), &model.UserDataType{
						ID:     id,
						PubKey: body.PubKey,
					})

					if err != nil {
						log.Println("DB Error:", err.Error())

						ctx.Status(500).Send("Internal Server Error")
					} else {
						token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
							Audience:  id,
							ExpiresAt: time.Now().UTC().Add(time.Minute * 15).Unix(),
						})

						tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

						if err != nil {
							log.Println("JWT Error:", err.Error())

							ctx.Status(500).Send("Internal Server Error")

							return
						}

						refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
							Audience: id,
							Subject:  "Refresh",
						})

						refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_KEY")))

						if err != nil {
							log.Println("JWT Error:", err.Error())

							ctx.Status(500).Send("Internal Server Error")

							return
						}

						ctx.JSON(&model.RegistrationResponse{
							PubKey:       body.PubKey,
							ID:           id,
							Token:        tokenString,
							RefreshToken: refreshTokenString,
						})
					}
				} else {
					ctx.Status(409).Send("Aleady Registered")
				}

			} else {
				ctx.Status(400).Send("Bad Request")
			}
		}
	}
}
