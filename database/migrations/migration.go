package migrations

import (
	"github.com/goblog/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Migrate() {
	db, err := gorm.Open("mysql", "root:ab789789@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	DB = db
	DB.DropTableIfExists(&models.User{}, &models.Article{}, &models.Passport{}, &models.City{}, &models.PetCategory{}, &models.ArticleCategory{}, &models.ArticleComment{})
	DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Passport{}, &models.City{}, &models.PetCategory{}, &models.ArticleCategory{}, &models.ArticleComment{})
	defer DB.Close()

}
