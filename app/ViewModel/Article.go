package ViewModel

import (
	"github.com/goblog/app/models"
)

type ArticleVM struct {
	models.BaseModel
	Title               string                 `json:"title"`
	Content             string                 `json:"content"`
	UserID              int                    `json:"user_id"`
	PetCategoryID       int                    `json:"pet_category_id"`
	PetCategory         models.PetCategory     `json:"pet_category"`
	ArticleCategoryID   int                    `json:"article_category_id"`
	ArticleCategory     models.ArticleCategory `json:"article_category"`
	Ctr                 int                    `json:"ctr"`
	ImgUrl              string                 `json:"imgurl"`
	Name                string                 `json"name"`
	Account             string                 `json:"account" `
	ArticleCommentCount int                    `json:"article_comment_count"`
}
