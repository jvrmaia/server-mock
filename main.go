package main

import (
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	config "github.com/jvrmaia/server-mock/config"
	handlers "github.com/jvrmaia/server-mock/handlers"
	"github.com/jvrmaia/server-mock/logger"
)

func main() {
	log := logger.New(false)
	var server config.Routes

	if _, err := toml.DecodeFile("config.toml", &server); err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("starting routes configuration")
	h := mux.NewRouter()
	for _, r := range server.Routes {
		switch r.Type {
		case "generic":
			h.HandleFunc(r.Path, handlers.GenericHandler(r.StatusCode, r.Path, r.ContentType, r.Headers.ToMap(), r.Body, server.Debug))
		case "echo":
			h.HandleFunc(r.Path, handlers.EchoHandler)
		default:
			h.HandleFunc(r.Path, handlers.GenericHandler(r.StatusCode, r.Path, r.ContentType, r.Headers.ToMap(), r.Body, server.Debug))
		}

		log.Info().Msg("route added: " + r.Path)
	}
	log.Info().Msg("routes configuration end")

	srv := &http.Server{
		Handler:      h,
		Addr:         server.Listen,
		WriteTimeout: server.WriteTimeout.Duration,
		ReadTimeout:  server.ReadTimeout.Duration,
		IdleTimeout:  server.IdleTimeout.Duration,
	}

	log.Info().Msg("starting server")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server startup failed")
	}
}
