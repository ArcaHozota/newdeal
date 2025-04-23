package service

import (
	"context"
	"errors"
	"newdeal/common"
	"newdeal/common/tools"
	"newdeal/ent/student"
	"time"
)

type LoginRequest struct {
	Username string `form:"username"`
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
	checkPass := tools.CheckHashPassword(loginUser.Password, loginForm.Password)
	if !checkPass {
		return common.EmptyString, errors.New(common.PasswordError)
	}
	return loginUser.LoginAccount, nil
}
