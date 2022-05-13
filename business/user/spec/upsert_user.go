package spec

type UpsertUserCreateSpec struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UpsertUserUpdateSpec struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Role  int    `validate:"required"`
}

type UpsertLoginUserSpec struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UpsertUserPasswordUpdateSpec struct {
	OldPassword     string `validate:"required"`
	NewPassword     string `validate:"required"`
	ConfirmPassword string `validate:"required,eqfield=NewPassword"`
}

type UpsertUserForgotPasswordSpec struct {
	Email string `validate:"required,email"`
}

type UpsertUserResetPasswordSpec struct {
	NewPassword     string `validate:"required"`
	ConfirmPassword string `validate:"required,eqfield=NewPassword"`
}
