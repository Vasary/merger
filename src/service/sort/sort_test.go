package sort

import (
	"merger/src/entity"
	"testing"
)

var dataProvider = []struct {
	points     int64
	trustLevel int64
}{
	{-1, 100},
	{3, 100},
	{5, 100},
	{5, 100},
	{3, 100},
	{15, 100},
	{12, 100},
	{12, 100},
	{15, 100},
	{15, 100},
	{15, 100},
	{15, 300},
	{15, 200},
}

func TestSort(t *testing.T) {
	var list []entity.User

	for _, pointEntity := range dataProvider {
		var user entity.User
		user.IncreasePoints(pointEntity.points)
		user.SetTrustLevel(pointEntity.trustLevel)

		list = append(list, user)
	}

	Sort(list)

	if list[0].GetPoints() != 15 && list[0].GetTrustLevel() == 300 {
		t.Error("Sorting error")
	}
}
