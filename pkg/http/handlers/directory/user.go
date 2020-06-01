package directory

import (
	"context"
	"strings"

	"github.com/gargakshit/paperplane-server/database"
	"github.com/gargakshit/paperplane-server/model"
	"github.com/gofiber/fiber"
)

// GetUserPublicKey is the handler to return the user's public key
func GetUserPublicKey(ctx *fiber.Ctx) {
	id := strings.ToUpper(ctx.Params("id"))

	directoryCollection := database.MongoConnection.Database("paperplane").Collection("directory")

	var resultUser model.UserDataType

	if err := directoryCollection.FindOne(context.TODO(), &model.UserDataType{
		ID: id,
	}).Decode(&resultUser); err != nil {
		ctx.Status(404).Send("Not found")
	} else {
		ctx.JSON(resultUser)
	}
}
