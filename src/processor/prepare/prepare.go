package prepare

import (
	"merger/src/entity"
	"merger/src/service/user/detail"
	"merger/src/logger"
	"github.com/jinzhu/gorm"
)

func Prepare(list entity.UserList, db *gorm.DB) []entity.User {
	logger.Info("Preparing row data")

	var users []entity.User

	for _, userId := range list.GetIdList() {
		users = append(users, detail.GetUserDetail(userId, db))
	}

	return users
}
