package seeders

import (
	"github.com/goblog/database/factories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Seed() {
	db, err := gorm.Open("mysql", "root:ab789789@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	DB = db
	UserSeeder()
	ArticleSeeder()
	CitySeeder()
	ArticleCategorySeeder()
	PetCategorySeeder()
	ArticleCommentSeeder()

}
func UserSeeder() {
	for i := 1; i <= 50; i++ {
		DB.Create(factory.UserFake())

	}

}
func ArticleSeeder() {
	for i := 1; i <= 100000; i++ {
		DB.Create(factory.ArticleFake())
	}

}
func CitySeeder() {
	cities := factory.CityFake()
	for _, city := range cities {
		DB.Create(&city)

	}
}
func ArticleCategorySeeder() {
	articleCategories := factory.ArticleCategoryFake()
	for _, articleCategory := range articleCategories {
		DB.Create(&articleCategory)

	}
}
func PetCategorySeeder() {
	petCategories := factory.PetCategoryFake()
	for _, petCategory := range petCategories {
		DB.Create(&petCategory)

	}
}
func ArticleCommentSeeder() {
	for i := 1; i <= 200; i++ {
		DB.Create(factory.ArticleCommentFake())
	}
}
