package sort

import (
	"merger/entity"
	"github.com/bradfitz/slice"
)

func Sort(users []entity.User) []entity.User {
	slice.Sort(users[:], func(a, b int) bool {
		return users[a].GetPoints() < users[b].GetPoints()
	})

	return users
}
