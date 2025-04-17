package repository

import (
	"newdeal/common"
	"newdeal/models"
)

func CountHymns() (uint32, error) {
	var kennsu int64
	err := models.DB.Where(&models.Hymn{VisibleFlg: true}).Count(&kennsu).Error
	return uint32(kennsu), err
}

func PaginationHymns(keyword string, pageSize uint8, offset uint32) ([]models.Hymn, error) {
	searchStr := common.GetDetailKeyword1(keyword)
	var hymns []models.Hymn
	err := models.DB.Model(&models.Hymn{}).InnerJoins("inner joins hymns_work on hymns_work.work_id = hymns.id").Where(&models.Hymn{NameJp: searchStr}).Or(&models.Hymn{NameKr: searchStr}).Or(&models.HymnWork{NameJpRa: searchStr}).Limit(int(pageSize)).Offset(int(offset)).Find(&hymns).Error
	return hymns, err
}
