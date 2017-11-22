package repositories

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/models"
	"strconv"
)

type ArticleRepository struct{}

func (ArticleRepository ArticleRepository) GetArticle(c *gin.Context) interface{} {
	query := DB
	articles := []models.Article{}
	// articles1 := []models.Article{}

	pageQuery := c.DefaultQuery("page", "1")
	per_pageQuery, _ := strconv.ParseFloat(c.DefaultQuery("per_page", "5"), 64)
	petCategoryQuery := c.DefaultQuery("pet_category_id", "0")
	page, _ := strconv.ParseFloat(pageQuery, 64)
	ArticleCategoryQuery := c.DefaultQuery("article_category_id", "0")

	CtrQuery := c.DefaultQuery("ctr", "0")

	CreateQuery := c.DefaultQuery("created_at", "desc")

	if petCategoryQuery != "0" {
		query = query.Where("pet_category_id = ?", petCategoryQuery)
	}
	if ArticleCategoryQuery != "0" {
		query = query.Where("article_category_id = ?", ArticleCategoryQuery)
	}
	if CtrQuery == "desc" {
		query = query.Order("ctr DESC")
	} else if CtrQuery == "asc" {
		query = query.Order("ctr ASC")
	}
	if CreateQuery != "desc" {
		query = query.Order("created_at ASC")
	} else {
		query = query.Order("created_at DESC")

	}
	query2 := query
	query = query.Select("articles.*,count(distinct ac.id) as article_comment_count").Joins("left join article_comments as ac on articles.id=ac.article_id").Group("articles.id")
	// DB.Model(&users).Related(&articles, "Article")
	// DB.Offset(5).Limit(10).Find(&articles)
	result := PaginateWithCount(models.ArticleTable(), query, &articles, per_pageQuery, page, query2)
	// result := Paginate(models.ArticleTable(), query, &articles, 5, page)

	return result
}
