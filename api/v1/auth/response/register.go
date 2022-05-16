package response

type RegisterResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}

func CreateRegisterResponse(name string, email string, verifyCode string) *RegisterResponse {
	return &RegisterResponse{
		Name:       name,
		Email:      email,
		VerifyCode: verifyCode,
	}
}
