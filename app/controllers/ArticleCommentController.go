package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/ViewModel"
	"github.com/goblog/app/models"
	"github.com/goblog/app/repositories"
	mapper "github.com/jeevatkm/go-model"
)

type ArticleCommentController struct {
	BaseController
	ArticleCommentRepository repositories.ArticleCommentRepository
}

func (con ArticleCommentController) init() {
	con.ArticleCommentRepository = repositories.ArticleCommentRepository{}

}

func (con ArticleCommentController) Index(c *gin.Context) {

	//	c.JSON(200, gin.H{
	//		"data": con.ArticleCommentRepository.GetArticleComment(),
	//	})
}

func (con ArticleCommentController) Create(c *gin.Context) {
	ArticleCommentPostVM := &ViewModel.ArticleCommentPost{}
	Binderr := c.BindJSON(ArticleCommentPostVM)
	Comment := &models.ArticleComment{}
	if Binderr == nil {
		mapper.Copy(Comment, ArticleCommentPostVM)
		Comment.Ip = c.ClientIP()
		con.ArticleCommentRepository.CreateArticleComment(Comment)
		c.JSON(200, gin.H{
			"status": "success",
		})
	} else {
		c.JSON(404, gin.H{
			"status": "error",
		})
	}

}
