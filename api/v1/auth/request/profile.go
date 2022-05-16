package request

import "vidlearn-final-projcect/business/user/spec"

type UpdateProfileRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func (request *UpdateProfileRequest) ToSpec(role int) *spec.UpsertUserUpdateSpec {
	return &spec.UpsertUserUpdateSpec{
		Name:  request.Name,
		Email: request.Email,
		Role:  role,
	}
}
