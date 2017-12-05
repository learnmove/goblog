package controllers

// github.com/goblog/controllers/user
import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/ViewModel"
	"github.com/goblog/app/models"
	"github.com/goblog/app/repositories"
	mapper "github.com/jeevatkm/go-model"
	"strings"
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
func (con ArticleController) GetArticles(c *gin.Context) {
	articles := con.ArticleRepository.GetArticles2(c.Copy())

	c.JSON(200, gin.H{
		"data": articles,
	})
}
func (con ArticleController) GetArticle(c *gin.Context) {
	article := con.ArticleRepository.GetArticle(c.Copy())
	ArticleVM := ViewModel.ArticleWithComment{}
	mapper.Copy(&ArticleVM, article)
	c.JSON(200, gin.H{
		"data": ArticleVM,
	})
}
func (con ArticleController) PostArticle(c *gin.Context) {
	articleVM := &ViewModel.ArticlePost{}
	bindErr := c.BindJSON(articleVM)
	articleVM.Content = strings.Replace(articleVM.Content, "\n", "<br>", -1)
	article := &models.Article{}
	mapper.Copy(article, articleVM)

	if bindErr != nil {
		c.JSON(403, gin.H{
			"error": "綁定錯誤",
		})
	} else {
		CreateErr := con.ArticleRepository.PostArticle(article)
		if CreateErr != nil {
			c.JSON(403, gin.H{
				"error": "無法新增",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "success",
			})
		}
	}

}
func (con ArticleController) Test(c *gin.Context) {

	c.JSON(403, gin.H{
		"status": con.ArticleRepository.GetArticles2(c.Copy()),
	})
}
