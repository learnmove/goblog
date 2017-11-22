package controllers

// github.com/goblog/controllers/user
import (
	// "fmt"
	"github.com/gin-gonic/gin"
	// "github.com/goblog/app/ViewModel"
	"github.com/goblog/app/repositories"
	// mapper "github.com/jeevatkm/go-model"
)

type ArticleController struct {
	BaseController
	UserRepository    repositories.UserRepository
	ArticleRepository repositories.ArticleRepository
}

func (con ArticleController) init() {
	con.UserRepository = repositories.UserRepository{}
	con.ArticleRepository = repositories.ArticleRepository{}
}
func (con ArticleController) GetArticle(c *gin.Context) {
	articles := con.ArticleRepository.GetArticle(c.Copy())

	c.JSON(200, gin.H{
		"articles": articles,
	})
}
