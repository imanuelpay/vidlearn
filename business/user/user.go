package user

import "time"

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Role       int
	IsReset    int
	VerifyCode string
	VerifiedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUser(
	name string,
	email string,
	password string,
	verifyCode string,
	createdAt time.Time,
) *User {
	return &User{
		Name:       name,
		Email:      email,
		Password:   password,
		Role:       0,
		VerifyCode: verifyCode,
		CreatedAt:  createdAt,
		VerifiedAt: time.Time{},
	}
}

func (oldUser *User) ModifyUser(
	name string,
	email string,
	role int,
	updatedAt time.Time,
) *User {
	return &User{
		ID:         oldUser.ID,
		Name:       name,
		Email:      email,
		Role:       role,
		Password:   oldUser.Password,
		VerifyCode: oldUser.VerifyCode,
		VerifiedAt: oldUser.VerifiedAt,
		UpdatedAt:  updatedAt,
	}
}

func (oldUser *User) ModifyPassword(
	password string,
) *User {
	return &User{
		ID:         oldUser.ID,
		Name:       oldUser.Name,
		Email:      oldUser.Email,
		Role:       oldUser.Role,
		Password:   password,
		VerifyCode: oldUser.VerifyCode,
		VerifiedAt: oldUser.VerifiedAt,
		UpdatedAt:  time.Now(),
	}
}

func (oldUser *User) ModifyVerify(
	verifyCode string,
	verifiedAt time.Time,
) *User {
	return &User{
		ID:         oldUser.ID,
		Name:       oldUser.Name,
		Email:      oldUser.Email,
		Role:       oldUser.Role,
		Password:   oldUser.Password,
		VerifyCode: verifyCode,
		VerifiedAt: verifiedAt,
		UpdatedAt:  time.Now(),
	}
}

func (oldUser *User) ModifyVerifyCode(
	verifyCode string,
) *User {
	return &User{
		ID:         oldUser.ID,
		Name:       oldUser.Name,
		Email:      oldUser.Email,
		Role:       oldUser.Role,
		Password:   oldUser.Password,
		VerifyCode: verifyCode,
		VerifiedAt: oldUser.VerifiedAt,
		UpdatedAt:  time.Now(),
	}
}
