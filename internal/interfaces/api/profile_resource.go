package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func HandlerFromMux(pr ProfileResource, r chi.Router) http.Handler {
	r.Route("/profiles", func(r chi.Router) {
		r.Get("/", pr.GetProfiles)
		r.Get("/{userId}", pr.GetProfileByUserId)
		r.Post("/", pr.CreateProfile)
	})

	return r
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
// @Router /api/profiles [get]
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
// @Router /api/profiles/{userId} [get]
func (pr ProfileResource) GetProfileByUserId(w http.ResponseWriter, r *http.Request) {
	userIdParam := chi.URLParam(r, "userId")
	userId, err := uuid.Parse(userIdParam)
	if err != nil {
		render.Render(w, r, responses.ErrorRenderer(fmt.Errorf("invalid format for parameter userId: %v", err)))
		return
	}

	profile := responses.NewProfileResponse(
		userId,
		"test",
		"test",
	)

	render.Respond(w, r, profile)
}

// CreateProfile godoc
// @Summary Create profile
// @Description create profile
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param profile body requests.CreateProfileRequest true "Create profile"
// @Success 201 {object} responses.ProfileResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/profiles [post]
func (pr ProfileResource) CreateProfile(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		render.Render(w, r, responses.ErrUnsupportedMediaType)
		return
	}

	request := &requests.CreateProfileRequest{}
	if err := render.Decode(r, request); err != nil {
		render.Render(w, r, responses.ErrorRenderer(err))
		return
	}

	// err := h.app.Commands.CreateProfileCommandHandler.Handle(r.Context(), command)
	// if err != nil {
	// 	util.WriteResponse(w, "Failed to create profile "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}
