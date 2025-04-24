package service

import (
	"context"
	"errors"
	"newdeal/common"
	"newdeal/common/tools"
	"newdeal/ent"
	"newdeal/ent/student"
	"time"
)

type LoginRequest struct {
	Username string `form:"loginAcct"`
	Password string `form:"password"`
}

func ProcessLogin(loginForm LoginRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	username := loginForm.Username
	loginUser, err := EntCore.Student.Query().
		Where(
			student.VisibleFlg(true),
			student.Or(
				student.LoginAccountEQ(username),
				student.EmailEQ(username),
			),
		).Only(ctx)
	if err != nil {
		return common.EmptyString, errors.New(common.StudentError)
	}
	err = preLogin(*loginUser)
	if err != nil {
		return common.EmptyString, err
	}
	checkPass := tools.CheckHashPassword(loginUser.Password, loginForm.Password)
	if !checkPass {
		return common.EmptyString, errors.New(common.PasswordError)
	}
	return loginUser.LoginAccount, nil
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
