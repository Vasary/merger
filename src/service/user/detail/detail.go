package detail

import (
	"merger/src/entity"
	"merger/src/logger"
	"merger/src/service/database"
	"encoding/json"
	"time"
	"fmt"
)

func GetUserDetail(id int64) entity.User {
	row := database.GetConnection().Raw("SELECT COALESCE(identity_mask, '{}'::json) as identity_mask, id, created_at, updated_at FROM users WHERE id = ?", id).Row()

	var mask string
	var userId int64
	var createdAt time.Time
	var updatedAt time.Time

	row.Scan(&mask, &userId, &createdAt, &updatedAt)

	var dat map[string]string

	if err := json.Unmarshal([]byte(mask), &dat); err != nil {
		logger.Warning(fmt.Sprintf("%d users mask not found", id))
	} else {
		logger.Info(fmt.Sprintf("Mask for user %d is %v", id, mask))
	}

	var maskEntity entity.UserMask
	var user entity.User

	maskEntity.Name = getValue(dat, "name")
	maskEntity.Surname = getValue(dat, "surname")
	maskEntity.Patronymic = getValue(dat, "patronymic")
	maskEntity.Birthday = getValue(dat, "birthday")

	user.Id = userId
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	user.Mask = maskEntity

	return user
}

func getValue(dat map[string]string, key string) string {
	value, ok := dat[key]
	if ok {
		return value
	}

	return ""
}
