package response

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func CreateLoginResponse(name string, email string, token string) *LoginResponse {
	return &LoginResponse{
		Name:  name,
		Email: email,
		Token: token,
	}
}
