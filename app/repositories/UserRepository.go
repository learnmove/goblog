package repositories

import (
	"github.com/goblog/app/models"
)

type UserRepository struct{}

func (UserRepository UserRepository) GetUser() []models.User {
	// var articles ArticleRepository
	users := []models.User{}
	// DB.Model(&users).Related(&articles, "Article")
	// DB.Find(&users)
	DB.Preload("Articles").Find(&users)
	DB.Model(&users).Preload("Articles")

	// for i, _ := range users {
	// 	DB.Model(users[i]).Related(&users[i].Articles)
	// }
	return users

}
