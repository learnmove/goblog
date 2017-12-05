package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/repositories"
)

type CategoryController struct {
	BaseController
	CategoryRepository repositories.CategoryRepository
}

func (con CategoryController) init() {
	con.CategoryRepository = repositories.CategoryRepository{}

}

func (con CategoryController) GetPetAndCityAndCategory(c *gin.Context) {
	cities := con.CategoryRepository.GetCityCategory()
	petCategories := con.CategoryRepository.GetPetCategory()
	articleCategory := con.CategoryRepository.GetArticleategory()
	c.JSON(200, gin.H{
		"data": gin.H{
			"cities":   cities,
			"pets":     petCategories,
			"category": articleCategory,
		},
	})
}
