package db

import (
	"fmt"

	"github.com/fondazionecrocereale/navegayaapi/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "my-db-cluster-do-user-17449050-0.i.db.ondigitalocean.com"
	port     = 25060
	user     = "doadmin"
	password = "AVNS_H9tD5yIYs9uct631SZJ"
	dbName   = "defaultdb"
)

func DBconnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db

}
