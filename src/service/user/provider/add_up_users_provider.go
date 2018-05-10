package provider

import (
	"merger/src/entity"
	"merger/src/logger"
	"merger/src/handler"
	"merger/src/service/database"
	"github.com/lib/pq"
)

func GetUsers() []entity.UserList {
	logger.Info("Looking for users for add up")

	rows, err := database.GetConnection().Raw("select up.value, array_agg(user_id) from user_parameters up join parameter p on up.parameter_id = p.id and p.id = 5 group by up.value, source_id having count(up.id) > 1").Rows()

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