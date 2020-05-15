package directory

import (
	"log"
	"strings"

	"github.com/gargakshit/paperplane-server/database"
	"github.com/gargakshit/paperplane-server/model"
	"github.com/gofiber/fiber"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// GetUserPublicKey is the handler to return the user's public key
func GetUserPublicKey(ctx *fiber.Ctx) {
	id := strings.ToUpper(ctx.Params("id"))

	res, err := r.Table("directory").GetAllByIndex("id", id).Run(database.RethinkSession)
	defer res.Close()

	if err != nil {
		log.Println("DB Error:", err.Error())

		ctx.Status(500).Send("Internal Server Error")
	} else {
		if res.IsNil() {
			ctx.Status(404).Send("Not found")
		} else {
			var user *model.UserDataType
			res.One(&user)
			ctx.Send(user.PubKey)
		}
	}
}
