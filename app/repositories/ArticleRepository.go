package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/models"
	"github.com/jinzhu/gorm"

	"strconv"
)

type ArticleRepository struct{}

func (ArticleRepository ArticleRepository) GetArticles(c *gin.Context) interface{} {
	query := DB
	articles := []models.Article{}
	// articles1 := []models.Article{}

	pageQuery := c.DefaultQuery("page", "1")
	per_pageQuery, _ := strconv.ParseFloat(c.DefaultQuery("per_page", "15"), 64)
	petCategoryQuery := c.DefaultQuery("pet_category_id", "0")
	CityQuery := c.DefaultQuery("city_id", "0")
	page, _ := strconv.ParseFloat(pageQuery, 64)
	ArticleCategoryQuery := c.DefaultQuery("article_category_id", "0")

	CtrQuery := c.DefaultQuery("ctr", "0")

	CreateQuery := c.DefaultQuery("created_at", "0")

	if petCategoryQuery != "0" {
		query = query.Where("pet_category_id = ?", petCategoryQuery)
	}
	if ArticleCategoryQuery != "0" {
		query = query.Where("article_category_id = ?", ArticleCategoryQuery)
	}
	if CityQuery != "0" {
		query = query.Where("city_id = ?", CityQuery)
	}

	query2 := query
	if CreateQuery == "desc" || CreateQuery == "0" {
		query = query.Order("articles.created_at desc")

	} else {
		query = query.Order("articles.created_at asc")
	}
	if CtrQuery == "desc" {
		query = query.Order("articles.ctr DESC")
	} else if CtrQuery == "asc" {
		query = query.Order("articles.ctr ASC")
	}
	query = query.Select("articles.id,articles.created_at,articles.title,articles.pet_category_name,articles.article_category_name,articles.city_name,articles.ctr,articles.img_url,articles.name,count(ac.id) as article_comment_count").Joins("left join article_comments as ac on articles.id=ac.article_id").Group("articles.id")

	result := PaginateWithCount(models.ArticleTable(), query, &articles, per_pageQuery, page, query2)
	return result
}
func (ArticleRepository ArticleRepository) GetArticles2(c *gin.Context) interface{} {
	whereCaluse := ""
	andControl := 0
	articles := &[]models.Article{}
	per_pageQuery, _ := strconv.ParseFloat(c.DefaultQuery("per_page", "15"), 64)
	page, _ := strconv.ParseFloat(c.DefaultQuery("page", "1"), 64)
	petCategoryQuery := c.DefaultQuery("pet_category_id", "0")
	CityQuery := c.DefaultQuery("city_id", "0")
	ArticleCategoryQuery := c.DefaultQuery("article_category_id", "0")
	if petCategoryQuery != "0" || ArticleCategoryQuery != "0" || CityQuery != "0" {
		whereCaluse = "where "
		if petCategoryQuery != "0" {
			if andControl == 0 {
				andControl = andControl + 1
			}
			whereCaluse = whereCaluse + "pet_category_id = " + petCategoryQuery
		}
		if ArticleCategoryQuery != "0" {
			if andControl == 0 {
				andControl = andControl + 1
			} else {
				whereCaluse = whereCaluse + " and "
			}
			whereCaluse = whereCaluse + "article_category_id = " + ArticleCategoryQuery
		}
		if CityQuery != "0" {
			if andControl == 0 {
				andControl = andControl + 1
			} else {
				whereCaluse = whereCaluse + " and "
			}
			whereCaluse = whereCaluse + "city_id = " + CityQuery
		}

	}
	queryCaluse := `select a.*  ,count(b.id) as article_comment_count  from articles as a  left  JOIN (select id,article_id from article_comments) as b on a.id=b.article_id  inner join (select id from articles  ` + whereCaluse + `  order by created_at desc limit offset, take) as c on a.id=c.id  Group BY a.id  order by created_at desc `
	countCaluse := `select count(*) as total from articles ` + whereCaluse
	result := PaginateFast(queryCaluse, countCaluse, articles, per_pageQuery, page)
	return result
}
func (ArticleRepository ArticleRepository) GetArticle(c *gin.Context) *models.Article {
	id := c.Params.ByName("id")
	query := DB
	article := &models.Article{}
	query.Preload("ArticleComments", func(db *gorm.DB) *gorm.DB {
		return db.Select("created_at,user_id,name,account,img_url,content,id")
	}).Where("id=?", id).Find(article)
	return article

}
func (ArticleRepository ArticleRepository) PostArticle(article *models.Article) error {
	return DB.Create(article).Error

}

func (ArticleRepository ArticleRepository) GetTest() interface{} {
	article := &[]models.Article{}
	DB.Raw(`
   select a.*  ,count(b.id) from articles as a  left  JOIN (select id,article_id from article_comments) as b on a.id=b.article_id  inner join (select id from articles limit 155000,511) as c on a.id=c.id   Group BY a.id  ORDER BY a.created_at asc`).Scan(article)

	return article

}
