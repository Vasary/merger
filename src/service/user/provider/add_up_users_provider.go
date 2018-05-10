package provider

import (
	"merger/src/entity"
	"merger/src/logger"
	"merger/src/handler"
	"github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

func GetUsers(db *gorm.DB) []entity.UserList {
	logger.Info("Looking for users for add up")

	rows, err := db.Raw("select up.value, array_agg(user_id) from user_parameters up join parameter p on up.parameter_id = p.id and p.id = ? group by up.value, source_id having count(up.id) > 1", 5).Rows()

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