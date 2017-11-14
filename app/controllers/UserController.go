package controllers

// github.com/goblog/controllers/user
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/models"
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
func (con UserController) Create(c *gin.Context) {
	user := &models.User{}
	fmt.Println(string(user.Password))
	c.BindJSON(user)
	registerErr := con.UserRepository.Register(user)
	if registerErr != nil {
		c.JSON(403, gin.H{
			"error": "sorry it fault register",
		})
	} else {
		c.JSON(200, gin.H{
			"name": "ok",
		})
	}

}
func (con UserController) Login(c *gin.Context) {
	user := &models.User{}
	c.BindJSON(user)
	LoginUser, result, token := con.UserRepository.Login(user)
	if result == false {
		c.JSON(403, gin.H{
			"error": "登錄失敗",
		})
	} else {
		c.JSON(200, gin.H{
			"user":  LoginUser,
			"token": token,
		})
	}

	// if LoginUser != nil {
	// 	c.JSON(200, gin.H{
	// 		"name": LoginUser.Password,
	// 	})

	// } else {
	// 	c.JSON(403, gin.H{
	// 		"error": "sorry not user",
	// 	})
	// }

}
func (con UserController) GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"users":    con.UserRepository.GetUser(),
		"articles": con.ArticleRepository.GetArticle(),
	})
}
func (con UserController) Test(c *gin.Context) {
	abc := c.Param("abc")
	fmt.Println(c.Request.URL.Path + "***************")
	c.HTML(http.StatusOK, "main/index.tmpl",
		gin.H{
			"abc": abc,
		})
}
