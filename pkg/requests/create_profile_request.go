package requests

import "github.com/google/uuid"

type CreateProfileRequest struct {
	UserId    uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}
