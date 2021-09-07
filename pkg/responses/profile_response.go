package responses

import "github.com/google/uuid"

type ProfileResponse struct {
	UserId    uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func NewProfileResponse(
	userId uuid.UUID,
	firstName, lastName string) ProfileResponse {
	return ProfileResponse{
		UserId:    userId,
		FirstName: firstName,
		LastName:  lastName,
	}
}
