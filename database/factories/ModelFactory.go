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
	cityID := rand.Intn(21) + 1
	city := CityFake()[cityID].Name
	categoryId := rand.Intn(16) + 1
	category := ArticleCategoryFake()[categoryId].Name
	petcategoryId := rand.Intn(9) + 1
	petcategory := PetCategoryFake()[petcategoryId].Name
	article := Article{
		Title:               fake.Title(),
		Content:             fake.Paragraph(),
		UserID:              rand.Intn(50) + 1,
		Ctr:                 rand.Intn(100) + 1,
		Ip:                  fake.IPv4(),
		Name:                fake.FirstName(),
		Account:             fake.FirstName(),
		CityName:            city,
		CityID:              cityID + 1,
		PetCategoryID:       petcategoryId + 1,
		PetCategoryName:     petcategory,
		ArticleCategoryID:   categoryId + 1,
		ArticleCategoryName: category,
	}

	return &article
}
func CityFake() []City {
	cities := []City{
		{Name: "基隆市"},
		{Name: "嘉義市"},
		{Name: "台北市"},
		{Name: "嘉義縣"},
		{Name: "新北市"},
		{Name: "台南市"},
		{Name: "桃園縣"},
		{Name: "高雄市"},
		{Name: "新竹市"},
		{Name: "新竹縣"},
		{Name: "台東縣"},
		{Name: "苗栗縣"},
		{Name: "屏東縣"},
		{Name: "花蓮縣"},
		{Name: "台中市"},
		{Name: "宜蘭縣"},
		{Name: "彰化縣"},
		{Name: "澎湖縣"},
		{Name: "南投縣"},
		{Name: "金門縣"},
		{Name: "雲林縣"},
		{Name: "連江縣"},
	}
	return cities
}
func ArticleCategoryFake() []ArticleCategory {
	articleCategory := []ArticleCategory{
		{Name: "綜合"},
		{Name: "救援"},
		{Name: "搞笑"},
		{Name: "活動"},
		{Name: "聚會"},
		{Name: "獸醫"},
		{Name: "志工"},
		{Name: "疾病"},
		{Name: "中途"},
		{Name: "認養"},
		{Name: "教學"},
		{Name: "祝褔"},
		{Name: "回憶"},
		{Name: "販售"},
		{Name: "捐贈/免費"},
		{Name: "內幕/八卦"},
		{Name: "其它"},
	}
	return articleCategory
}
func PetCategoryFake() []PetCategory {
	petCategory := []PetCategory{
		{Name: "狗"},
		{Name: "貓"},
		{Name: "倉鼠"},
		{Name: "兔子"},
		{Name: "鳥"},
		{Name: "昆蟲"},
		{Name: "蠍子"},
		{Name: "蜥蜴"},
		{Name: "蜘蛛"},
		{Name: "蛇"},
	}
	return petCategory
}
func ArticleCommentFake() interface{} {
	articleComment := ArticleComment{
		Content:   fake.Paragraph(),
		UserID:    rand.Intn(50) + 1,
		ArticleID: rand.Intn(50) + 1,
		Ip:        fake.IPv4(),
		Name:      fake.FirstName(),
		Account:   fake.FirstName(),
	}

	return &articleComment
}
