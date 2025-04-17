package repository

import (
	"newdeal/common"
	"newdeal/models"
)

func CountHymns() uint32 {
	var kennsu int64
	models.DB.Where(&models.Hymn{VisibleFlg: true}).Count(&kennsu)
	return uint32(kennsu)
}

func PaginationHymns(keyword string, pageSize uint8, offset uint32) []models.Hymn {
	searchStr := common.GetDetailKeyword1(keyword)
	var hymns []models.Hymn
	models.DB.Model(&models.Hymn{}).InnerJoins("inner joins hymns_work on hymns_work.work_id = hymns.id").Where(&models.Hymn{NameJp: searchStr}).Or(&models.Hymn{NameKr: searchStr}).Or(&models.HymnWork{NameJpRa: searchStr}).Limit(int(pageSize)).Offset(int(offset)).Find(&hymns)
	return hymns
}
