package main

import (
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {
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
