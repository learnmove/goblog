package repositories

import (
	"fmt"
	"github.com/goblog/app/ViewModel"
	"github.com/goblog/app/helper"
	"github.com/goblog/app/models"
)

type UserRepository struct{}

func (UserRepository UserRepository) GetUser() []models.User {
	// var articles ArticleRepository
	users := []models.User{}
	// DB.Model(&users).Related(&articles, "Article")
	// DB.Find(&users)
	DB.Preload("Articles").Find(&users)
	// DB.Model(&users).Preload("Articles")

	// for i, _ := range users {
	// 	DB.Model(users[i]).Related(&users[i].Articles)
	// }
	return users

}
func (UserRepository *UserRepository) Register(user *models.User) error {
	hash, hasherr := helper.HashPassword(user.Password)
	if hasherr != nil {
		fmt.Println("hash error")
		return hasherr
	} else {
		user.Password = string(hash)
		return DB.Create(user).Error
	}
}
func (UserRepository *UserRepository) Login(user *ViewModel.UserLoginPost) (models.User, bool, string) {
	var t string = ""
	LoginUser := models.User{}
	DB.Where("account = ?", user.Account).Find(&LoginUser)
	if LoginUser.ID == 0 {
		return LoginUser, false, t
	}
	// compare password
	fmt.Println(user)
	CompareErr := helper.ComparePassword(LoginUser.Password, user.Password)
	if CompareErr != true {
		fmt.Println("error password")
		return LoginUser, false, t

	} else {
		// sucess login and write token in db

		t = helper.GenerateToken()
		// DB.Create(&models.Passport{Account: LoginUser.Account, Token: t})
		DB.Where(models.Passport{Account: LoginUser.Account}).Assign(models.Passport{Token: t}).FirstOrCreate(&models.Passport{})

		return LoginUser, true, t
	}

}
