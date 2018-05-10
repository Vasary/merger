package sort

import (
	"merger/src/entity"
	"testing"
)

var dataProvider = []struct {
	points int64
}{
	{-1},
	{3},
	{5},
	{5},
	{3},
	{15},
	{12},
	{12},
	{15},
	{15},
	{15},
	{15},
	{15},
}

func TestSort(t *testing.T) {
	var list []entity.User

	for _, pointEntity := range dataProvider {
		var user entity.User
		user.IncreasePoints(pointEntity.points)

		list = append(list, user)
	}

	Sort(list)

	if list[0].GetPoints() != 15 {
		t.Error("Sorting error")
	}
}
