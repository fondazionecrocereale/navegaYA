package db

import (
	"fmt"

	"github.com/fondazionecrocereale/navegayaapi/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "private-db-postgresql-nyc3-68472-do-user-17449050-0.m.db.ondigitalocean.com"
	port     = 25060
	user     = "doadmin"
	contra = "AVNS_0RVpY_d_VjNV9QF5sTe"
	dbName   = "defaultdb"
)

func DBconnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, contra, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db

}
