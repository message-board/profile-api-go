package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/message-board/profile-go/internal/interfaces/api"
)

func main() {
	var addr string
	pflag.StringVarP(&addr, "address", "a", ":8080", "the address for the api to listen on. Host and port separated by ':'")
	pflag.Parse()

	// gracefully exit on keyboard interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// configure logger
	log, _ := zap.NewProduction(zap.WithCaller(false))
	defer func() {
		_ = log.Sync()
	}()

	srv := api.NewServer(log, addr)
	go func() {
		if err := srv.RunServer(func(router chi.Router) http.Handler {
			return api.HandlerFromMux(
				api.NewProfileResource(),
				router,
			)
		}); err != nil {
			os.Exit(1)
		}
	}()

	<-c
	log.Info("gracefully shutting down")
	os.Exit(0)
}
