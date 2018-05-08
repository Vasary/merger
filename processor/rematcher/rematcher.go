package rematcher

import (
	"usersRematcher/entity"
	"usersRematcher/logger"
	"usersRematcher/service/database"
	"usersRematcher/handler"
	"fmt"
)

func Rematch(users []entity.User) {
	logger.Info.Println("Rematching")

	logger.Info.Println(fmt.Sprintf("Main user id %d", users[0].Id))

	for i := 1; i < len(users); i++ {
		logger.Info.Println( fmt.Sprintf("Mergin %d with %d", users[i].Id, users[0].Id))

		_, err := database.GetConnection().Raw(`SELECT rematch_import(?, ?)`, users[0].Id, users[i].Id).Rows()

		deleteUser(users[i].Id)

		handler.FailOnError(err, "Rematching error")
	}
}

func deleteUser(userId int64) {

}
