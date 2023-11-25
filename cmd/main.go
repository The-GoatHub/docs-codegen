package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"

	"The-GoatHub/docs-codegen/internal/api"
)

const defaultPort = 3000

func main() {
	var (
		ctx  = context.Background()
		mux  = chi.NewRouter()
		quit = make(chan os.Signal, 1)
	)

	fmt.Println(uuid.New())

	api.HandlerFromMux(
		api.NewStrictHandler(api.NewServer(), nil),
		mux)
	api.RegisterStatic(mux)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3000),
		Handler: mux,
	}
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info(ctx, "server started", "port", defaultPort)
		if err := server.ListenAndServe(); err != nil {
			log.Error(ctx, "starting http server", "err", err)
		}
	}()

	<-quit
}
