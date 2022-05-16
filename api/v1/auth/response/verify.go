package response

import "time"

type VerifyResponse struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	VerifiedAt time.Time `json:"verified_at"`
}

func CreateVerifyResponse(name string, email string, verifiedAt time.Time) *VerifyResponse {
	return &VerifyResponse{
		Name:       name,
		Email:      email,
		VerifiedAt: verifiedAt,
	}
}
