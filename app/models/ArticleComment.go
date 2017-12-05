package models

// . "github.com/goblog/app/models/article"

type ArticleComment struct {
	BaseModel
	Content   string `json:"content" binding:"required" gorm:"not null"`
	UserID    int    `json:"user_id" binding:"required" gorm:"not null"`
	ArticleID int    `json:"article_id" binding:"required" gorm:"not null;index:idx_article_id"`
	Ip        string `json:"ip"`
	ImgUrl    string `json:"img_url"`
	User      User   `json:"-"`
	Name      string `json":name" binding:"required gorm:"not null"`
	Account   string `json:"account" binding:"required gorm:"not null"`
}

func ArticleCommentTable() string {
	return "article_comments"
}
