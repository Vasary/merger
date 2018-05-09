package provider

import (
	"merger/src/entity"
	"merger/src/logger"
	"merger/src/handler"
	"merger/src/service/database"
	"github.com/lib/pq"
)

func GetUsers() []entity.UserList {
	logger.Info.Println("Looking for users for addup")

	// rows, err := Database.GetConnection().Raw("select name as value, array_agg(name) as list from parameter group by name").Rows()
	rows, err := database.GetConnection().Raw("select up.value, array_agg(user_id) from user_parameters up join parameter p on up.parameter_id = p.id and p.id = 5 where user_id IN (38323041,38326365,38326639,38326368,38322863,38322857) group by up.value, source_id having count(up.id) > 1").Rows()

	handler.FailOnError(err, "Error in query")

	defer rows.Close()

	var data []entity.UserList

	for rows.Next() {
		var value string
		var ar []int64
		err = rows.Scan(&value, pq.Array(&ar))

		data = append(data, entity.UserList{Value: value, List: ar})
	}

	return data
}