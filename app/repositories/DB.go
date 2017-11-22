package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {

	db, err := gorm.Open("mysql", "root:ab789789@/goblog?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		return
	}
	DB = db
}
