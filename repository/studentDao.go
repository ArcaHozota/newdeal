package repository

import "newdeal/models"

func FindStudentsByEmail(email string) []models.Student {
	var students []models.Student
	models.DB.Where("visible_flg = true AND (login_account = ? OR email = ?)", email, email).Find(&students)
	return students
}

func GetStudentById(id int64) []models.Student {
	var student []models.Student
	models.DB.First(&student, models.Student{VisibleFlg: true, Id: id})
	return student
}

func CountStudentsByAccount(account string, id int64) uint32 {
	var kennsu int64
	models.DB.Where(&models.Student{VisibleFlg: true, LoginAccount: account}).Not(&models.Student{Id: id}).Count(&kennsu)
	return uint32(kennsu)
}
