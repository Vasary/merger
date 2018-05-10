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

	logger.Info(fmt.Sprintf("Master user id is %d", users[0].Id))

	for i := 1; i < len(users); i++ {
		logger.Info( fmt.Sprintf("Merging %d with %d", users[i].Id, users[0].Id))

		_, err := database.GetConnection().Raw(`SELECT merge_users(?, ?)`, users[0].Id, users[i].Id).Rows()

		handler.FailOnError(err, "Rematching error")
	}

	logger.Info("Row merged")
}