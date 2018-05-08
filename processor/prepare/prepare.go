package prepare

import (
	"usersRematcher/entity"
	"usersRematcher/service/user/detail"
	"usersRematcher/logger"
)

func Prepare(list entity.UserList) []entity.User {
	logger.Info.Println("Prepearing row data")

	var users []entity.User

	for _, userId := range list.GetIdList() {
		users = append(users, detail.GetUserDetail(userId))
	}

	return users
}
