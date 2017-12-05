package ViewModel

import (
	"github.com/goblog/app/models"
)

type ArticleComment struct {
	models.BaseModel
	Content string `json:"content"`
	ImgUrl  string `json:"imgurl"`
	Name    string `json"name" `
	Account string `json:"account"`
}
