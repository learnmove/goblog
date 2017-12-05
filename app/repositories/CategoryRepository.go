package repositories

import (
	"github.com/goblog/app/models"
)

type CategoryRepository struct{}

func (CategoryRepository CategoryRepository) GetCityCategory() []models.City {

	cities := []models.City{}

	DB.Find(&cities)
	return cities
}
func (CategoryRepository CategoryRepository) GetPetCategory() []models.PetCategory {

	petCategories := []models.PetCategory{}

	DB.Find(&petCategories)
	return petCategories
}
func (CategoryRepository CategoryRepository) GetArticleategory() []models.ArticleCategory {

	articleCategory := []models.ArticleCategory{}

	DB.Find(&articleCategory)
	return articleCategory
}
