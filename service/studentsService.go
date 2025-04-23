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
	Username string `json:"loginAcct" binding:"required"`
	Password string `json:"password" binding:"required"`
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
		return common.EmptyString, err
	}
	checkPass := tools.CheckHashPassword(loginUser.Password, loginForm.Password)
	if !checkPass {
		return common.EmptyString, errors.New("password error")
	}
	return loginUser.LoginAccount, nil
}
