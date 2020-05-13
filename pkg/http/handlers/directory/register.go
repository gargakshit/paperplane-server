package directory

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gargakshit/paperplane-server/database"
	"github.com/gargakshit/paperplane-server/model"
	"github.com/gargakshit/paperplane-server/utils"
	"github.com/gofiber/fiber"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// RegisterHandler is the HTTP handler for the /register endpoint
func RegisterHandler(ctx *fiber.Ctx) {
	var body model.RegistrationBody
	err := ctx.BodyParser(&body)

	if err != nil {
		log.Panicln("HTTP Error:", err.Error())

		ctx.Status(500).Send("Internal Server Error")
	} else {
		if err != nil {
			log.Println("HTTP Error:", err.Error())

			ctx.Status(500).Send("Internal Server Error")
		} else {
			if utils.IsBase64Valid(body.PubKey) {
				id := utils.GenerateRandomID()

				res, err := r.Table("directory").Get(body.PubKey).Run(database.RethinkSession)

				defer res.Close()

				if err != nil {
					log.Println("DB Error:", err.Error())

					ctx.Status(500).Send("Internal Server Error")
				}

				if res.IsNil() {
					err = r.Table("directory").Insert(model.UserDataType{
						ID:     id,
						PubKey: body.PubKey,
					}).Exec(database.RethinkSession)

					if err != nil {
						log.Println("DB Error:", err.Error())

						ctx.Status(500).Send("Internal Server Error")
					} else {
						token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
							Audience:  fmt.Sprintf("reg--%s", id),
							ExpiresAt: time.Now().UTC().Add(time.Minute * 15).Unix(),
						})

						tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

						if err != nil {
							log.Println("JWT Error:", err.Error())

							ctx.Status(500).Send("Internal Server Error")
						}

						ctx.JSON(&model.RegistrationResponse{
							PubKey: body.PubKey,
							ID:     id,
							Token:  tokenString,
						})
					}
				} else {
					ctx.Status(409).Send("Already Registered")
				}

			} else {
				ctx.Status(400).Send("Bad Request")
			}
		}
	}
}
