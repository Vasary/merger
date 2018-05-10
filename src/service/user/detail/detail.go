package detail

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"merger/src/entity"
	"merger/src/handler"
	"merger/src/logger"
)

func GetUserDetail(id int64, db *gorm.DB) entity.User {
	row := db.Raw(`
		SELECT
			COALESCE(u.identity_mask, '{}' :: json) as identity_mask,
  			u.id,
  			s.trust_level
		
		FROM users u
  			JOIN user_parameters up ON up.user_id = u.id AND up.parameter_id = 5
  			JOIN source s ON s.id = up.source_id

		WHERE u.id = ?
	`, id)

	var mask string
	var userId int64
	var trustLevel int64

	row.Row().Scan(&mask, &userId, &trustLevel)

	if userId == 0 {
		handler.Fail(fmt.Sprintf("Requested user id %v, obtained %v", id, userId))
	}

	var dat map[string]string

	if err := json.Unmarshal([]byte(mask), &dat); err != nil {
		logger.Debug(fmt.Sprintf("%d users mask not found", id))
	} else {
		logger.Debug(fmt.Sprintf("Mask for user %d is %v", id, mask))
	}

	var maskEntity entity.UserMask
	var user entity.User

	maskEntity.Name = getValue(dat, "name")
	maskEntity.Surname = getValue(dat, "surname")
	maskEntity.Patronymic = getValue(dat, "patronymic")
	maskEntity.Birthday = getValue(dat, "birthday")

	user.SetId(userId)
	user.SetMask(maskEntity)
	user.SetTrustLevel(trustLevel)

	return user
}

func getValue(dat map[string]string, key string) string {
	value, ok := dat[key]
	if ok {
		return value
	}

	return ""
}
