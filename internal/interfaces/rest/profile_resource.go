package rest

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/message-board/profile-go/pkg/requests"
	"github.com/message-board/profile-go/pkg/responses"
)

type ProfileResource struct {
}

func NewProfileResource() ProfileResource {
	return ProfileResource{}
}

// ListProfiles godoc
// @Summary List profiles
// @Description get profiles
// @Tags profiles
// @Accept  json
// @Produce  json
// @Success 200 {array} responses.ProfileResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/users [get]
func (pr ProfileResource) GetProfiles(w http.ResponseWriter, r *http.Request) {
	profile := []responses.ProfileResponse{
		responses.NewProfileResponse(uuid.New(), "test1", "test1"),
		responses.NewProfileResponse(uuid.New(), "test2", "test2"),
	}
	render.Respond(w, r, profile)
}

// GetProfile godoc
// @Summary Get profile
// @Description get profile by user id
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param userId path int true "User ID"
// @Success 200 {object} responses.ProfileResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/users/{id} [get]
func (pr ProfileResource) GetProfile(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	profile := responses.NewProfileResponse(
		userId,
		"test",
		"test",
	)

	render.Respond(w, r, profile)
}

// CreateUser godoc
// @Summary Create user
// @Description create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body requests.CreateUserRequest true "Create user"
// @Success 201 {object} responses.ProfileResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/users [post]
func (pr ProfileResource) CreateProfile(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		render.Render(w, r, responses.ErrUnsupportedMediaType())
		return
	}

	request := &requests.CreateProfileRequest{}
	if err := render.Decode(r, request); err != nil {
		render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	// err := h.app.Commands.CreateUserCommandHandler.Handle(r.Context(), command)
	// if err != nil {
	// 	util.WriteResponse(w, "Failed to create user "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}
