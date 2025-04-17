package repository

import (
	"database/sql"
	"newdeal/common"
	"newdeal/common/pagination"
	"newdeal/models"

	"gorm.io/gorm"
)

func CountHymnsAll() (int, error) {
	var kennsu int64
	err := models.DB.Model(&models.Hymn{}).Where(&models.Hymn{VisibleFlg: true}).Count(&kennsu).Error
	return int(kennsu), err
}

func CountHymnsByKeyword(keyword string) (int, error) {
	searchStr := common.GetDetailKeyword1(keyword)
	var kennsu int64
	err := models.DB.Model(&models.Hymn{}).
		InnerJoins("INNER JOIN hymns_work ON hymns_work.work_id = hymns.id").
		Where("hymns.name_jp LIKE ? OR hymns.name_kr LIKE ? OR hymns_work.name_jp_rational LIKE ?", searchStr, searchStr, searchStr).
		Count(&kennsu).Error
	return int(kennsu), err
}

func PaginationHymns(keyword string, pageNum int, pageSize int) ([]models.Hymn, error) {
	searchStr := common.GetDetailKeyword1(keyword)
	var hymns []models.Hymn
	err := models.DB.
		Model(&models.Hymn{}).
		InnerJoins(`
			INNER JOIN hymns_work
			    ON hymns_work.work_id = hymns.id
		`).
		Scopes(ByKeyword(searchStr), pagination.PaginateScope(pageNum, pageSize)).
		Find(&hymns).Error
	return hymns, err
}

// ---------- scope_helpers.go ----------
func ByKeyword(kw string) func(*gorm.DB) *gorm.DB {
	like := common.GetDetailKeyword1(kw)
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(`
			hymns.name_jp LIKE @kw OR
			hymns.name_kr LIKE @kw OR
			hymns_work.name_jp_rational LIKE @kw
		`, sql.Named("kw", like))
	}
}
