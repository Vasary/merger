package prepare

import (
	"merger/entity"
	"merger/service/user/detail"
	"merger/logger"
)

func Prepare(list entity.UserList) []entity.User {
	logger.Info.Println("Prepearing row data")

	var users []entity.User

	for _, userId := range list.GetIdList() {
		users = append(users, detail.GetUserDetail(userId))
	}

	return users
}
