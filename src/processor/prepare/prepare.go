package prepare

import (
	"github.com/jinzhu/gorm"
	"merger/src/entity"
	"merger/src/logger"
	"merger/src/service/user/detail"
)

func Prepare(list entity.UserList, db *gorm.DB) []entity.User {
	logger.Info("Preparing row data")

	var users []entity.User

	for _, userId := range list.GetIdList() {
		users = append(users, detail.GetUserDetail(userId, db))
	}

	return users
}
