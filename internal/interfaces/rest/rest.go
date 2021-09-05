package rest

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RunServer(createHandler func(router chi.Router) http.Handler) {
	RunServerOnAddr(":"+os.Getenv("PORT"), createHandler)
}

func RunServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
	rootRouter := chi.NewRouter()

	// ref: https://github.com/swaggo/http-swagger
	rootRouter.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(addr+"/swagger/doc.json"),
	))

	logrus.Info("Starting HTTP server")

	http.ListenAndServe(addr, rootRouter)
}
