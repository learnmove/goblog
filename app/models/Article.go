package models

// . "github.com/goblog/app/models/article"

type Article struct {
	BaseModel
	Title               string           `json:"title"`
	Content             string           `json:"content"`
	UserID              int              `json:"user_id"`
	PetCategoryID       int              `json:"pet_category_id" gorm:"index:idx_pet"`
	PetCategoryName     string           `json:"pet_category_name"`
	ArticleCategoryID   int              `json:"article_category_id" gorm:"index:idx_cate"`
	ArticleCategoryName string           `json:"article_category_name"`
	CityID              int              `json:"city_id" gorm:"index:idx_city"`
	CityName            string           `json:"city_name"`
	Ctr                 int              `json:"ctr"`
	Ip                  string           `json:"ip"`
	ImgUrl              string           `json:"img_url"`
	ArticleComments     []ArticleComment `json:"article_comment" `
	User                User             `json:"-"`
	Name                string           `json:"name" binding:"required gorm:"not null"`
	Account             string           `json:"account" binding:"required gorm:"not null"`
	ArticleCommentCount int              `json:"article_comment_count" sql:"-"`
}

func ArticleTable() string {
	return "articles"
}
