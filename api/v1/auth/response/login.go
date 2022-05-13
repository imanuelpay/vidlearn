package response

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func CreateLoginResponse(id int, name string, email string, token string) *LoginResponse {
	return &LoginResponse{
		ID:    id,
		Name:  name,
		Email: email,
		Token: token,
	}
}
