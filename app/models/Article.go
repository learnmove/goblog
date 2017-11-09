package models

// . "github.com/goblog/app/models/article"

type Article struct {
	BaseModel
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}
