package ViewModel

// . "github.com/goblog/app/models/article"

type ArticlePost struct {
	Title               string `json:"title"`
	Content             string `json:"content"`
	UserID              int    `json:"user_id"`
	PetCategoryID       int    `json:"pet_category_id"`
	PetCategoryName     string `json:"pet_category_name"`
	ArticleCategoryID   int    `json:"article_category_id"`
	ArticleCategoryName string `json:"article_category_name"`
	CityID              int    `json:"city_id"`
	CityName            string `json:"city_name"`
	ImgUrl              string `json:"img_url"`
	Name                string `json:"name" binding:"required"`
	Account             string `json:"account" binding:"required"`
}
