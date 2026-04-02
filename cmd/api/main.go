package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {

	//root context
	ctx := context.Background()

	// initialize the app
	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer func() {
		if err := app.Close(ctx); err != nil {
			log.Printf("failed to close app resources: %v", err)
		}
	}()
	router := httpserver.NewRouter()
	// standard go type that runs the http server
	srv := &http.Server{
		Addr:              ":5000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("API running on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
