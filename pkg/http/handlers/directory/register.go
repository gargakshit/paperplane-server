package directory

import (
	"log"

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
						ctx.JSON(&model.RegistrationResponse{
							PubKey: body.PubKey,
							ID:     id,
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
