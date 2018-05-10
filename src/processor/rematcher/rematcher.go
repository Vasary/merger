package rematcher

import (
	"merger/src/entity"
	"merger/src/logger"
	"merger/src/service/database"
	"merger/src/handler"
	"fmt"
)

func Rematch(users []entity.User) {
	logger.Info("Rematching")

	logger.Info(fmt.Sprintf("Main user id %d", users[0].Id))

	for i := 1; i < len(users); i++ {
		logger.Info( fmt.Sprintf("Mergin %d with %d", users[i].Id, users[0].Id))

		_, err := database.GetConnection().Raw(`SELECT rematch_import(?, ?)`, users[0].Id, users[i].Id).Rows()

		deleteUser(users[i].Id)

		handler.FailOnError(err, "Rematching error")
	}
}

func deleteUser(userId int64) {

}
