package models

type ArticleCategory struct {
	BaseModel
	Name string `json:"name"  gorm:"not null"`
}
