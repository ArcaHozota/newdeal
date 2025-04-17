package repository

import "newdeal/models"

func CountHymns() int64 {
	var kennsu int64
	models.DB.Where(&models.Hymn{VisibleFlg: true}).Count(&kennsu)
	return kennsu
}
