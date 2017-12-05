package migrations

import (
	"github.com/goblog/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateIndex() {
	db, err := gorm.Open("mysql", "root:ab789789@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	DB = db
	// article table
	db.Model(&models.Article{}).AddIndex("idx_p_a_c_c", "pet_category_id", "article_category_id", "city_id", "created_at")
	db.Model(&models.Article{}).AddIndex("idx_p_c_c", "pet_category_id", "city_id", "created_at")
	db.Model(&models.Article{}).AddIndex("idx_p_a_c", "pet_category_id", "article_category_id", "created_at")
	db.Model(&models.Article{}).AddIndex("idx_a_c_c", "article_category_id", "city_id", "created_at")
	db.Model(&models.Article{}).AddIndex("idx_p_c", "pet_category_id", "created_at")
	db.Model(&models.Article{}).AddIndex("idx_a_c", "article_category_id", "created_at")
	db.Model(&models.Article{}).AddIndex("idx_c_c", "city_id", "created_at")
	// no created_at for count
	db.Model(&models.Article{}).AddIndex("idx_p__a", "pet_category_id", "article_category_id")
	db.Model(&models.Article{}).AddIndex("idx_p__c", "pet_category_id", "city_id")
	db.Model(&models.Article{}).AddIndex("idx_a__c", "article_category_id", "city_id")
	db.Model(&models.Article{}).AddIndex("idx_p__a__c", "pet_category_id", "article_category_id", "city_id")
	defer DB.Close()

}
