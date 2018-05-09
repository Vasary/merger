package score

import (
	"merger/src/entity"
	"merger/src/logger"
	"merger/src/service/validator/birthday"
	"merger/src/service/validator/name"
	"merger/src/service/validator/surname"
	"merger/src/service/validator/patronymic"
	"merger/src/service/sort"
)

func Score(users []entity.User) []entity.User {
	logger.Info.Println("Scoring user mask")

	for index, user := range users {
		users[index].IncreasePoints(score(user))
	}

	return sort.Sort(users)
}

func score(user entity.User) int64 {
	var points int64

	if birthday.Validate(user.GetMask().Birthday) {
		points += 12
	}

	if name.Validate(user.GetMask().Name) {
		points += 10
	}

	if surname.Validate(user.GetMask().Surname) {
		points += 11
	}

	if patronymic.Validate(user.GetMask().Patronymic) {
		points += 8
	}

	return points
}