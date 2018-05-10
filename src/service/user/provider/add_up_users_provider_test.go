package provider

import (
	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"testing"
	"fmt"
)

var DB *gorm.DB

var commonReply = []map[string]interface{}{
	{"value": "one", "array_agg": "{1,2,3,4}"},
	{"value": "two", "array_agg": "{4,5,6}"},
}

func TestGetUsers(t *testing.T)  {
	setupMock()

	GlobalMock := mocket.Catcher
	GlobalMock.Logging = false
	GlobalMock.NewMock().WithArgs(int64(5)).WithReply(commonReply)

	users := GetUsers(DB)

	t.Run("Rows for duplicate", func(t *testing.T) {
		if len(users) != 2 {
			t.Error(fmt.Sprintf("Expected 2 got %v", len(users)))
		}
	})

	t.Run("Data", func(t *testing.T) {
		for index, user := range users {
			if commonReply[index]["value"] != user.GetValue() {
				t.Error(fmt.Sprintf("Expected %v got %v", commonReply[index]["value"], user.GetValue()))
			}
		}
	})
}

func setupMock() {
	mocket.Catcher.Register()
	db, _ := gorm.Open(mocket.DRIVER_NAME, "connection string")

	DB = db
}