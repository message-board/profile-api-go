package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/message-board/profile-go/api/openapi" // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

// @title Message Board Profile Api
// @version 1.0
// @description TODO.
// @termsOfService http://todo.io/terms

// @contact.name TODO
// @contact.url http://todo.io/support
// @contact.email support@todo.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1

type Server struct {
	log  *zap.Logger
	addr string
}

func NewServer(log *zap.Logger, addr string) Server {
	return Server{
		log:  log,
		addr: addr,
	}
}

func (s Server) RunServer(createHandler func(router chi.Router) http.Handler) error {
	router := chi.NewRouter()
	s.setMiddlewares(router)

	router.Get("/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// we are mounting all APIs under /api path
	router.Mount("/api", createHandler(router))

	s.log.Info("Starting HTTP server")
	if err := http.ListenAndServe(s.addr, router); err != nil {
		s.log.Error("failed to start server", zap.Error(err))
		return err
	}

	s.log.Info("ready to serve requests on " + s.addr)
	return nil
}

func (s Server) setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}