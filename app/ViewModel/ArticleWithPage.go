package ViewModel

import (
	"github.com/goblog/app/models"
)

type ArticleWithPage struct {
	Data     ArticleVM    `json:"data"`
	PageInfo *models.Page `json:"pageinfo"`
}
