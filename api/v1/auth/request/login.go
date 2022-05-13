package request

import "vidlearn-final-projcect/business/user/spec"

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *LoginUserRequest) ToSpec() *spec.UpsertLoginUserSpec {
	return &spec.UpsertLoginUserSpec{
		Email:    request.Email,
		Password: request.Password,
	}
}
