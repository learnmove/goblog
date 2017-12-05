package repositories

import (
	"github.com/goblog/app/models"
)

type ArticleCommentRepository struct{}

func (ArticleCommentRepository ArticleCommentRepository) CreateArticleComment(comment *models.ArticleComment) error {

	return DB.Create(comment).Error
}
