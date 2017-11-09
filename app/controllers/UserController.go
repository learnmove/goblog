package controllers

// github.com/goblog/controllers/user
import (
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/repositories"
	"net/http"
)

type UserController struct {
	BaseController
	UserRepository    repositories.UserRepository
	ArticleRepository repositories.ArticleRepository
}

func (con UserController) init() {
	con.UserRepository = repositories.UserRepository{}
	con.ArticleRepository = repositories.ArticleRepository{}

}
func (con UserController) GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"users":    con.UserRepository.GetUser(),
		"articles": con.ArticleRepository.GetArticle(),
	})
}
func (con UserController) Test(c *gin.Context) {
	abc := c.Param("abc")

	c.HTML(http.StatusOK, "main/index.tmpl",
		gin.H{
			"abc": abc,
		})
}
