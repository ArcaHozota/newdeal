package service

import (
	"context"
	"errors"
	"newdeal/common"
	"newdeal/common/tools"
	"newdeal/ent"
	"newdeal/ent/student"
	"newdeal/pojos"
	"reflect"
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
		DateOfBirth:  loginUser.DateOfBirth.Format(common.DateLayout),
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
		DateOfBirth:  loginUser.DateOfBirth.Format(common.DateLayout),
	}, nil
}

func StudentInfoUpdate(studentDto pojos.StudentDTO) (string, error) {
	studentId, err := strconv.Atoi(studentDto.ID)
	if err != nil {
		return common.EmptyString, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	studentById, err := EntCore.Student.Query().
		Where(
			student.VisibleFlg(true),
			student.ID(int64(studentId)),
		).Only(ctx)
	if err != nil {
		return common.EmptyString, err
	}
	hikakuStudentDto := pojos.StudentDTO{
		ID:           strconv.Itoa(int(studentById.ID)),
		LoginAccount: studentById.LoginAccount,
		Username:     studentById.Username,
		DateOfBirth:  studentById.DateOfBirth.Format(common.DateLayout),
		Email:        studentById.Email,
		Password:     studentDto.Password,
	}
	if reflect.DeepEqual(hikakuStudentDto, studentDto) {
		if studentById.Password == studentDto.Password || tools.CheckHashPassword(studentById.Password, studentDto.Password) {
			return common.NochangeMsg, nil
		}
		password, err := tools.GenerateHashPassword(studentDto.Password)
		if err != nil {
			return common.EmptyString, err
		}
		err = EntCore.Student.UpdateOneID(studentById.ID).
			SetPassword(password).
			Exec(ctx)
		if err != nil {
			return common.EmptyString, err
		}
		return common.UpdatedMsg, nil
	}
	password, err := tools.GenerateHashPassword(studentDto.Password)
	if err != nil {
		return common.EmptyString, err
	}
	birthday, err := time.Parse(common.DateLayout, studentDto.DateOfBirth)
	if err != nil {
		return common.EmptyString, err
	}
	err = EntCore.Student.UpdateOneID(studentById.ID).
		SetLoginAccount(studentDto.LoginAccount).
		SetPassword(password).
		SetUsername(studentDto.Username).
		SetEmail(studentDto.Email).
		SetDateOfBirth(birthday).
		SetUpdatedTime(time.Now()).
		Where(
			student.VisibleFlg(true),
		).Exec(ctx)
	if err != nil {
		return common.EmptyString, err
	}
	return common.UpdatedMsg, nil
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
