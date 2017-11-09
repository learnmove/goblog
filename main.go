package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/controllers"
)

func main() {

	UserController := controllers.UserController{}
	ArticleController := controllers.ArticleController{}
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/user", UserController.GetUser)
	r.GET("/article", ArticleController.GetArticle)
	r.GET("/test/:abc", UserController.Test)
	r.Run(":8089")

}
