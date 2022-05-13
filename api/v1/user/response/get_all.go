package response

import (
	"time"
	"vidlearn-final-projcect/business/user"
)

type GetUserByIDResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Role       int       `json:"role"`
	VerifiedAt time.Time `json:"verified_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func CreateGetAllUserResponse(users []*user.User) []*GetUserByIDResponse {
	var usersResponse []*GetUserByIDResponse

	for _, user := range users {
		usersResponse = append(usersResponse, CreateGetUserByIDResponse(user))
	}

	return usersResponse
}

func CreateGetUserByIDResponse(user *user.User) *GetUserByIDResponse {
	return &GetUserByIDResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		VerifiedAt: user.VerifiedAt,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
