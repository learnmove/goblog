package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
	BaseController
	//	TestRepository repositories.TestRepository

}

func (con TestController) init() {
	//	con.TestRepository = repositories.TestRepository{}

}

func (con TestController) Test(c *gin.Context) {
	abc := c.Param("abc")
	fmt.Println(c.Request.URL.Path + "***************")
	c.HTML(http.StatusOK, "main/index.tmpl",
		gin.H{
			"abc": abc,
		})
}
func (con TestController) Test1(c *gin.Context) {
	type Tform struct {
		Name string `  binding:"required"`
		Age  int
	}
	var tf Tform
	c.BindJSON(&tf)
	c.JSON(http.StatusOK, gin.H{
		"name": tf.Age,
		"age":  tf.Name,
	})
}
