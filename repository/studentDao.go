package repository

import "newdeal/models"

func findStudentsByEmail(email string) []models.Student {
	students := []models.Student{}
	models.DB.Where("visible_flg = true AND (login_account = ? OR email = ?)", email, email).Find(&students)
	return students
}

func getStudentById(id int64) []models.Student {
	student := []models.Student{}
	models.DB.First(&student, models.Student{VisibleFlg: true, Id: id})
	return student
}

func countStudentsByAccount(account string, id int64) int64 {
	var kensu int64
	models.DB.Where(&models.Student{VisibleFlg: true, LoginAccount: account}).Not(&models.Student{Id: id}).Count(&kensu)
	return kensu
}
