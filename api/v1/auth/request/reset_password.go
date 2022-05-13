package request

import "vidlearn-final-projcect/business/user/spec"

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (request *ForgotPasswordRequest) ToSpec() *spec.UpsertUserForgotPasswordSpec {
	return &spec.UpsertUserForgotPasswordSpec{
		Email: request.Email,
	}
}

func (request *ResetPasswordRequest) ToSpec() *spec.UpsertUserResetPasswordSpec {
	return &spec.UpsertUserResetPasswordSpec{
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}
