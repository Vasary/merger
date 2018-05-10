package rematcher

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"merger/src/entity"
	"merger/src/handler"
	"merger/src/logger"
)

func Rematch(users []entity.User, db *gorm.DB) {
	logger.Debug("Rematching")

	logger.Debug(fmt.Sprintf("Master user id is %d", users[0].GetId()))

	if len(users) < 2 {
		handler.Fail("For merging should have two or more users")
	}

	for i := 1; i < len(users); i++ {
		logger.Debug(fmt.Sprintf("Merging %d with %d", users[i].GetId(), users[0].GetId()))

		_, err := db.Raw("SELECT merge_users(?, ?)", users[0].GetId(), users[i].GetId()).Rows()

		handler.FailOnError(err, "Rematching error")
	}

	logger.Debug("Row merged")
}
