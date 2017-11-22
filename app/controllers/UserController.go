package controllers

// github.com/goblog/controllers/user
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/ViewModel"
	"github.com/goblog/app/models"
	"github.com/goblog/app/repositories"
	mapper "github.com/jeevatkm/go-model"
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
	bindErr := c.BindJSON(user)

	if bindErr != nil {
		c.JSON(403, gin.H{
			"error": "空白BODY",
		})
	} else {
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

}
func (con UserController) Login(c *gin.Context) {
	user := &ViewModel.UserLoginPost{}
	bindErr := c.BindJSON(user)

	if bindErr != nil {
		c.JSON(402, gin.H{
			"error": "登錄失敗",
		})
		return
	}

	LoginUser, result, token := con.UserRepository.Login(user)

	LoginVM := ViewModel.UserLogin{}
	err := mapper.Copy(&LoginVM, LoginUser)

	if err != nil {
		fmt.Println("err mapper")
	}
	if result == false {
		c.JSON(402, gin.H{
			"error": "登錄失敗",
		})

	} else {
		c.JSON(200, gin.H{
			"user": LoginVM,

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
	articles := con.ArticleRepository.GetArticle(c.Copy())
	c.JSON(200, gin.H{
		"users":    con.UserRepository.GetUser(),
		"articles": articles,
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
