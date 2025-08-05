package main

import (
	"flag"
	"fmt"
	"net/http"
	"personal_blog/backend/internal/app"
	"personal_blog/backend/internal/routes"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Go blog backend server port")
	flag.Parse()

	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	sqlDB, err := application.DB.DB()
	if err != nil {
		application.Logger.Fatalf("failed to get raw DB: %v", err)
	}
	defer sqlDB.Close()

	r := routes.SetUpRoutes(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Printf("🚀 Starting server on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil {
		application.Logger.Fatal(err)
	}
}
