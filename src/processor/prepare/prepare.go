package prepare

import (
	"merger/src/entity"
	"merger/src/service/user/detail"
	"merger/src/logger"
)

func Prepare(list entity.UserList) []entity.User {
	logger.Info("Preparing row data")

	var users []entity.User

	for _, userId := range list.GetIdList() {
		users = append(users, detail.GetUserDetail(userId))
	}

	return users
}
