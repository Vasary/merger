package provider

import (
	"fmt"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"testing"
)

var DB *gorm.DB

var commonReply = []map[string]interface{}{
	{"value": "one", "array_agg": "{1,2,3,4}"},
	{"value": "two", "array_agg": "{4,5,6}"},
}

func TestGetUsers(t *testing.T) {
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
}

func setupMock() {
	mocket.Catcher.Register()
	db, _ := gorm.Open(mocket.DRIVER_NAME, "connection string")

	DB = db
}
