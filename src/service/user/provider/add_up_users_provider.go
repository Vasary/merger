package provider

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"merger/src/entity"
	"merger/src/handler"
	"merger/src/logger"
)

func GetUsers(db *gorm.DB) []entity.UserList {
	logger.Debug("Looking for users for add up")

	rows, err := db.Raw("select up.value, array_agg(user_id) from user_parameters up join parameter p on up.parameter_id = p.id and p.id = ? where user_id IN (39415644,39416143,39415186,39414782,39414779,39416147,39415184) group by up.value, source_id having count(up.id) > 1", 5).Rows()

	handler.FailOnError(err, "Error in query")

	var data []entity.UserList

	for rows.Next() {
		var value string
		var ar []int64
		err = rows.Scan(&value, pq.Array(&ar))

		data = append(data, entity.UserList{Value: value, List: ar})
	}

	return data
}
