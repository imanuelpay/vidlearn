package response

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateRegisterResponse(name string, email string) *RegisterResponse {
	return &RegisterResponse{
		Name:  name,
		Email: email,
	}
}
