package response

type ForgotPasswordResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	ResetToken string `json:"reset_token"`
}

func CreateForgotPasswordResponse(name string, email string, verifyCode string) *ForgotPasswordResponse {
	return &ForgotPasswordResponse{
		Name:       name,
		Email:      email,
		ResetToken: verifyCode,
	}
}
