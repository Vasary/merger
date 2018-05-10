package score

import (
	"fmt"
	"merger/src/entity"
	"testing"
)

var dataProvider = []struct {
	name       string
	surname    string
	patronymic string
	birthday   string
	points     int64
}{
	{
		name:       "Ivan",
		surname:    "Ivanov",
		patronymic: "Ivanovich",
		birthday:   "15-12-1992",
		points:     41,
	},
	{
		name:       "Ivan",
		surname:    "",
		patronymic: "",
		birthday:   "",
		points:     10,
	},
	{
		name:       "",
		surname:    "Ivanov",
		patronymic: "",
		birthday:   "",
		points:     11,
	},
	{
		name:       "",
		surname:    "",
		patronymic: "Ivanovich",
		birthday:   "",
		points:     8,
	},
	{
		name:       "",
		surname:    "",
		patronymic: "",
		birthday:   "15-12-1992",
		points:     12,
	},
}

func TestScore(t *testing.T) {

	var users []entity.User

	for _, element := range dataProvider {
		var mask = entity.UserMask{
			Name:       element.name,
			Surname:    element.surname,
			Patronymic: element.patronymic,
			Birthday:   element.birthday,
		}

		var usr = entity.User{}
		usr.SetMask(mask)

		users = append(users, usr)
	}

	t.Run("Scoring test", func(t *testing.T) {

		users = Score(users)

		for index, usr := range users {
			if dataProvider[index].points != usr.GetPoints() {
				t.Error(fmt.Sprintf("Score error. Expected %v but got %v", dataProvider[index].points, usr.GetPoints()))
			}
		}

	})
}
