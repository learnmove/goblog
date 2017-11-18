package factory

import (
	"github.com/goblog/app/helper"
	. "github.com/goblog/app/models"
	"github.com/icrowley/fake"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UserFake() interface{} {

	t, _ := helper.HashPassword("secret")
	user := User{
		Name:     fake.FirstName(),
		Age:      rand.Intn(100),
		Password: string(t),
		Account:  fake.FirstName(),
	}
	return &user

}
func ArticleFake() interface{} {
	article := Article{
		Title:   fake.Title(),
		Content: fake.Paragraph(),
		UserID:  rand.Intn(50) + 1,
	}

	return &article
}
