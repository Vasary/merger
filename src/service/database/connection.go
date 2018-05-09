package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"merger/src/handler"
)

const (
	host     = "data.cfv4.com"
	port     = 5434
	user     = "testdata"
	password = "zTLsv2QeatnA45Ks2sJ7PGg5"
	dbname   = "userdata"
)

var connection *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", dsn)

	connection = db

	handler.FailOnError(err, "Can't connect to database")
}

func GetConnection() *gorm.DB {
	return connection
}