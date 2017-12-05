package repositories

import (
	// "fmt"
	"fmt"
	"github.com/goblog/app/models"
	"github.com/jinzhu/gorm"
	"math"
	"strconv"
	"strings"
)

// func GetCount(query *gorm.DB, model interface{}) float64 {
// 	var count float64
// 	query.Find(model)

// }
func Paginate(table string, query *gorm.DB, model interface{}, take float64, page float64) interface{} {

	pageinfo := &models.Page{}
	var offset float64
	offset = take * (page - 1)
	query.Table(table).Count(&pageinfo.Total)
	query.Table(table).Limit(take).Offset(offset).Find(model)
	pageinfo.LastPage = math.Ceil(pageinfo.Total / take)
	pageinfo.CurrentPage = page
	pageinfo.PerPage = take
	pageinfo.From = offset + 1
	if (offset + take) >= pageinfo.Total {
		pageinfo.To = pageinfo.Total
		fmt.Println("shit")

	} else {
		pageinfo.To = offset + take
		fmt.Println("no shit")

	}
	result := MergeModel(model, pageinfo)
	return result
}
func PaginateFast(queryCaluse string, countCaluse string, model interface{}, take float64, page float64) interface{} {
	pageinfo := &models.Page{}
	DB.Raw(countCaluse).Scan(pageinfo)
	var offset float64
	offset = take * (page - 1)
	getOffsetQuery := strings.Replace(queryCaluse, "offset", strconv.FormatFloat(offset, 'f', 0, 64), -1)
	getTakeQuery := strings.Replace(getOffsetQuery, "take", strconv.FormatFloat(take, 'f', 0, 64), -1)

	DB.Raw(getTakeQuery).Scan(model)

	pageinfo.LastPage = math.Ceil(pageinfo.Total / take)
	pageinfo.CurrentPage = page
	pageinfo.PerPage = take
	pageinfo.From = offset + 1
	if (offset + take) >= pageinfo.Total {
		pageinfo.To = pageinfo.Total
		fmt.Println("shit")

	} else {
		pageinfo.To = offset + take
		fmt.Println("no shit")

	}
	result := MergeModel(model, pageinfo)
	return result

}
func PaginateWithCount(table string, query *gorm.DB, model interface{}, take float64, page float64, query2 *gorm.DB) interface{} {

	pageinfo := &models.Page{}
	var offset float64
	offset = take * (page - 1)
	query2.Table(table).Count(&pageinfo.Total)
	query.Table(table).Limit(take).Offset(offset).Find(model)
	pageinfo.LastPage = math.Ceil(pageinfo.Total / take)
	pageinfo.CurrentPage = page
	pageinfo.PerPage = take
	pageinfo.From = offset + 1
	if (offset + take) >= pageinfo.Total {
		pageinfo.To = pageinfo.Total
		fmt.Println("shit")

	} else {
		pageinfo.To = offset + take
		fmt.Println("no shit")

	}
	result := MergeModel(model, pageinfo)
	return result
}
func MergeModel(data interface{}, pageinfo *models.Page) interface{} {
	return struct {
		Data     interface{}  `json:"data"`
		PageInfo *models.Page `json:"pageinfo"`
	}{
		Data:     data,
		PageInfo: pageinfo,
	}
}
