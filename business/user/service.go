package user

import (
	"errors"
	"fmt"
	"time"
	"vidlearn-final-projcect/business/mail"
	"vidlearn-final-projcect/business/user/spec"
	"vidlearn-final-projcect/config"
	"vidlearn-final-projcect/util"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllUser() (user []*User, err error)
	GetUserByID(ID int) (user *User, err error)
	GetUserByEmail(email string) (user *User, err error)
	GetUserByVerifyCode(verifyCode string) (user *User, err error)
	CreateUser(user *User) (*User, error)
	UpdateUser(userCurrent *User, IDCurrent int) (*User, error)
	DeleteUser(ID int) (user *User, err error)
}

type Service interface {
	GetAllUser() (user []*User, err error)
	GetUserByID(ID int) (user *User, err error)
	GetUserByEmailAndPassword(upsertUserLoginSpec *spec.UpsertLoginUserSpec) (user *User, err error)
	VerifyEmail(verifyCode string) (user *User, err error)
	ForgotPassword(email string) (user *User, err error)
	ResetPassword(upsertUserResetPasswordSpec *spec.UpsertUserResetPasswordSpec, verifyCode string) (*User, error)
	CreateUser(upsertUserSpec *spec.UpsertUserCreateSpec) (*User, error)
	UpdateUser(upsertUserSpec *spec.UpsertUserUpdateSpec, IDCurrent int) (*User, error)
	UpdateUserPassword(upsertUserSpec *spec.UpsertUserPasswordUpdateSpec, IDCurrent int) (*User, error)
	DeleteUser(ID int) (user *User, err error)
}

type userService struct {
	repository Repository
	config     *config.AppConfig
	validate   *validator.Validate
}

func CreateService(repository Repository, config *config.AppConfig) Service {
	return &userService{
		repository: repository,
		validate:   validator.New(),
		config:     config,
	}
}

func (service *userService) GetAllUser() (user []*User, err error) {
	return service.repository.GetAllUser()
}

func (service *userService) GetUserByID(ID int) (user *User, err error) {
	return service.repository.GetUserByID(ID)
}

func (service *userService) VerifyEmail(verifyCode string) (user *User, err error) {
	user, err = service.repository.GetUserByVerifyCode(verifyCode)
	if err != nil {
		return nil, err
	}

	if user.isReset == true {
		return nil, errors.New("User is reset")
	}

	user.VerifiedAt = time.Now()
	user.VerifyCode = " "

	return service.repository.UpdateUser(user, user.ID)
}

func (service *userService) ForgotPassword(email string) (user *User, err error) {
	user, err = service.repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	verifyCode := util.RandomString(64)

	user.VerifyCode = verifyCode
	user.isReset = true

	mailService := mail.CreateService(service.config)

	mail := new(mail.Mail)
	mail.To = user.Email
	mail.Subject = "Reset your password"
	mail.Body = "Reset your password by clicking this link: " + service.config.App.URL + "/api/v1/reset-password/" + verifyCode
	mail.From = service.config.Mail.Username

	_, err = mailService.SendMail(mail)
	if err != nil {
		return nil, err
	}

	return service.repository.UpdateUser(user, user.ID)
}

func (service *userService) ResetPassword(upsertUserResetPasswordSpec *spec.UpsertUserResetPasswordSpec, verifyCode string) (*User, error) {
	err := service.validate.Struct(upsertUserResetPasswordSpec)
	if err != nil {
		return nil, err
	}

	user, err := service.repository.GetUserByVerifyCode(verifyCode)
	if err != nil {
		return nil, err
	}

	if user.isReset == false {
		return nil, errors.New("User is not reset")
	}

	passwordHash, err := util.HashPassword(upsertUserResetPasswordSpec.NewPassword)
	if err != nil {
		return nil, err
	}

	user.Password = passwordHash

	return service.repository.UpdateUser(user, user.ID)
}

func (service *userService) GetUserByEmailAndPassword(upsertUserLoginSpec *spec.UpsertLoginUserSpec) (user *User, err error) {
	err = service.validate.Struct(upsertUserLoginSpec)
	if err != nil {
		return nil, err
	}

	user, err = service.repository.GetUserByEmail(upsertUserLoginSpec.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	time := time.Time{}
	if user.VerifiedAt == time {
		return nil, errors.New("user not verified")
	}

	isTrue, err := util.CheckPasswordHash(upsertUserLoginSpec.Password, user.Password)
	if !isTrue || err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (service *userService) CreateUser(upsertUserSpec *spec.UpsertUserCreateSpec) (*User, error) {
	err := service.validate.Struct(upsertUserSpec)
	if err != nil {
		return nil, err
	}

	password, err := util.HashPassword(upsertUserSpec.Password)
	if err != nil {
		return nil, err
	}

	verifyCode := util.RandomString(64)

	user := NewUser(
		upsertUserSpec.Name,
		upsertUserSpec.Email,
		password,
		verifyCode,
		time.Now(),
	)

	mailService := mail.CreateService(service.config)

	mail := new(mail.Mail)
	mail.To = user.Email
	mail.Subject = "Verify your email"
	mail.Body = "Verify your email by clicking this link: " + service.config.App.URL + "/api/v1/verify/" + verifyCode
	mail.From = service.config.Mail.Username

	fmt.Println(service.config.App.URL)
	_, err = mailService.SendMail(mail)
	if err != nil {
		return nil, err
	}

	return service.repository.CreateUser(user)
}

func (service *userService) UpdateUser(upsertUserSpec *spec.UpsertUserUpdateSpec, IDCurrent int) (*User, error) {
	err := service.validate.Struct(upsertUserSpec)
	if err != nil {
		return nil, err
	}

	userCurrent, err := service.repository.GetUserByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	user := userCurrent.ModifyUser(
		upsertUserSpec.Name,
		upsertUserSpec.Email,
		upsertUserSpec.Role,
		time.Now(),
	)

	return service.repository.UpdateUser(user, IDCurrent)
}

func (service *userService) UpdateUserPassword(upsertUserSpec *spec.UpsertUserPasswordUpdateSpec, IDCurrent int) (*User, error) {
	err := service.validate.Struct(upsertUserSpec)
	if err != nil {
		return nil, err
	}

	userCurrent, err := service.repository.GetUserByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	isTrue, err := util.CheckPasswordHash(upsertUserSpec.OldPassword, userCurrent.Password)
	if !isTrue || err != nil {
		return nil, errors.New("invalid old password")
	}

	password, err := util.HashPassword(upsertUserSpec.NewPassword)
	if err != nil {
		return nil, err
	}

	user := userCurrent.ModifyPassword(
		password,
	)

	return service.repository.UpdateUser(user, IDCurrent)
}

func (service *userService) DeleteUser(ID int) (user *User, err error) {
	return service.repository.DeleteUser(ID)
}
