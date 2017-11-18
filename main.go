package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/controllers"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders: []string{"Content-type"},
	}))
	api := r.Group("/api")
	{
		UserController := controllers.UserController{}
		ArticleController := controllers.ArticleController{}
		v1 := api.Group("/v1")
		{
			v1.GET("/user", UserController.GetUser)
			v1.POST("/user/register", UserController.Create)
			v1.POST("/user/login", UserController.Login)
			v1.GET("/article", ArticleController.GetArticle)

		}
		TestController := controllers.TestController{}
		t := api.Group("/t")
		{
			t.GET("/a/:abc", TestController.Test)
			t.POST("/t1", TestController.Test1)

		}

	}
	r.Run(":3001")

}
