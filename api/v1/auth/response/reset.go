package response

type ResetPasswordResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateResetPasswordResponse(name string, email string) *ResetPasswordResponse {
	return &ResetPasswordResponse{
		Name:  name,
		Email: email,
	}
}
