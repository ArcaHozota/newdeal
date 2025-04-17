package repository

import "newdeal/models"

func CountHymns() int64 {
	var kennsu int64
	models.DB.Where(&models.Hymn{VisibleFlg: true}).Count(&kennsu)
	return kennsu
}

func PaginationHymns(keyword string, pageSize, offset int64) []models.Hymn {
	searchStr := models.GetDetailKeyword1(keyword)
	var hymns []models.Hymn
	models.DB.Where(&models.Hymn{NameJp: searchStr}).Or(&models.Hymn{NameKr: searchStr})
}
