package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbConn *gorm.DB

func ConnectToDB(){
	
}