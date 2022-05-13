package request

import "vidlearn-final-projcect/business/user/spec"

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int    `json:"role"`
}

type UpdatePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (request *UpdateUserRequest) ToSpec() *spec.UpsertUserUpdateSpec {
	return &spec.UpsertUserUpdateSpec{
		Name:  request.Name,
		Email: request.Email,
		Role:  request.Role,
	}
}

func (request *UpdatePasswordRequest) ToSpec() *spec.UpsertUserPasswordUpdateSpec {
	return &spec.UpsertUserPasswordUpdateSpec{
		OldPassword:     request.OldPassword,
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}
