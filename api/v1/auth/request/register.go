package request

import "vidlearn-final-projcect/business/user/spec"

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *RegisterUserRequest) ToSpec() *spec.UpsertUserCreateSpec {
	return &spec.UpsertUserCreateSpec{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
