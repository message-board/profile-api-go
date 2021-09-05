package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/message-board/profile-go/internal/interfaces/rest"
)

func main() {
	rest.RunServer(func(router chi.Router) http.Handler {
		return rest.HandlerFromMux(
			rest.NewProfileResource(),
			router,
		)
	})
}
