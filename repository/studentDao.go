package repository

import "newdeal/models"

func FindStudentsByEmail(email string) ([]models.Student, error) {
	var students []models.Student
	err := models.DB.Where("visible_flg = true AND (login_account = ? OR email = ?)", email, email).Find(&students).Error
	return students, err
}

func GetStudentById(id int64) ([]models.Student, error) {
	var student []models.Student
	err := models.DB.First(&student, models.Student{VisibleFlg: true, Id: id}).Error
	return student, err
}

func CountStudentsByAccount(account string, id int64) (uint32, error) {
	var kennsu int64
	err := models.DB.Where(&models.Student{VisibleFlg: true, LoginAccount: account}).Not(&models.Student{Id: id}).Count(&kennsu).Error
	return uint32(kennsu), err
}
