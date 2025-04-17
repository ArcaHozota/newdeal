package repository

import "newdeal/models"

func findStudentsByEmail(email string) []models.Student {
	var students []models.Student
	models.DB.Where("visible_flg = true AND (login_account = ? OR email = ?)", email, email).Find(&students)
	return students
}

func getStudentById(id int64) []models.Student {
	var student []models.Student
	models.DB.First(&student, models.Student{VisibleFlg: true, Id: id})
	return student
}

func countStudentsByAccount(account string, id int64) int64 {
	var kennsu int64
	models.DB.Where(&models.Student{VisibleFlg: true, LoginAccount: account}).Not(&models.Student{Id: id}).Count(&kennsu)
	return kennsu
}
