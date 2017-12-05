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
		AllowHeaders: []string{"Content-type", "X-XSRF-TOKEN"},
	}))
	api := r.Group("/api")
	{
		UserController := controllers.UserController{}
		ArticleController := controllers.ArticleController{}
		CategoryController := controllers.CategoryController{}
		ArticleCommentController := controllers.ArticleCommentController{}

		v1 := api.Group("/v1")
		{
			v1.GET("/user", UserController.GetUser)
			v1.POST("/user/register", UserController.Create)
			v1.POST("/user/login", UserController.Login)
			v1.GET("/article", ArticleController.GetArticles)
			v1.POST("/article", ArticleController.PostArticle)
			v1.GET("/article/:id", ArticleController.GetArticle)
			v1.POST("/article/:id/comment", ArticleCommentController.Create)
			// v1.GET("/test/", ArticleController.Test)
			// v1.GET("/article/:id/comment", ArticleController.Test)

			v1.GET("/category", CategoryController.GetPetAndCityAndCategory)

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
