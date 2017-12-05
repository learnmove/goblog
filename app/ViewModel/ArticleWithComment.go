package ViewModel

import (
	"github.com/goblog/app/models"
)

type ArticleWithComment struct {
	models.BaseModel
	Title               string                  `json:"title"`
	Content             string                  `json:"content"`
	UserID              int                     `json:"user_id"`
	ImgUrl              string                  `json:"imgurl"`
	Name                string                  `json:"name"`
	Account             string                  `json:"account" `
	PetCategoryName     string                  `json:"pet_category_name"`
	ArticleCategoryName string                  `json:"article_category_name"`
	CityName            string                  `json:"city_name"`
	ArticleComments     []models.ArticleComment `json:"article_comment"`
}
