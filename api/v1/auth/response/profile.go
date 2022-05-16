package response

import (
	"time"
	"vidlearn-final-projcect/business/user"
)

type ProfileResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Role       int       `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	VerifiedAt time.Time `json:"verified_at"`
}

func CreateProfileResponse(user *user.User) *ProfileResponse {
	return &ProfileResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		VerifiedAt: user.VerifiedAt,
	}
}
