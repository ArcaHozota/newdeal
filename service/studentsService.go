package service

import (
	"context"
	"errors"
	"newdeal/common"
	"newdeal/common/tools"
	"newdeal/ent"
	"newdeal/ent/student"
	"newdeal/pojos"
	"strconv"
	"time"
)

type LoginRequest struct {
	Username string `form:"loginAcct"`
	Password string `form:"password"`
}

func GetStudentById(id int64) (pojos.StudentDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	loginUser, err := EntCore.Student.Query().
		Where(
			student.VisibleFlg(true),
			student.ID(id),
		).Only(ctx)
	if err != nil {
		return pojos.StudentDTO{}, errors.New(common.StudentError)
	}
	return pojos.StudentDTO{
		ID:           strconv.Itoa(int(loginUser.ID)),
		LoginAccount: loginUser.LoginAccount,
		Username:     loginUser.Username,
		Password:     loginUser.Password,
		Email:        loginUser.Email,
		DateOfBirth:  loginUser.DateOfBirth,
	}, nil
}

func ProcessLogin(loginForm LoginRequest) (pojos.StudentDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	username := loginForm.Username
	loginUser, err := EntCore.Student.Query().
		Where(
			student.VisibleFlg(true),
			student.Or(
				student.LoginAccount(username),
				student.Email(username),
			),
		).Only(ctx)
	if err != nil {
		return pojos.StudentDTO{}, errors.New(common.StudentError)
	}
	checkPass := tools.CheckHashPassword(loginUser.Password, loginForm.Password)
	if !checkPass {
		return pojos.StudentDTO{}, errors.New(common.PasswordError)
	}
	err = preLogin(*loginUser)
	if err != nil {
		return pojos.StudentDTO{}, err
	}
	return pojos.StudentDTO{
		ID:           strconv.Itoa(int(loginUser.ID)),
		LoginAccount: loginUser.LoginAccount,
		Username:     loginUser.Username,
		Password:     loginUser.Password,
		Email:        loginUser.Email,
		DateOfBirth:  loginUser.DateOfBirth,
	}, nil
}

func preLogin(loginUser ent.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := EntCore.Student.UpdateOneID(loginUser.ID).
		SetUpdatedTime(time.Now()).
		Where(student.VisibleFlg(true)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
