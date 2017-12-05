package models

type ArticleCategory struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"  gorm:"not null"`
}
