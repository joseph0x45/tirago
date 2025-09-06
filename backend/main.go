package main

import (
	"backend/handlers"
	"backend/middleware"
	"backend/store"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectToDB() (*sqlx.DB, error) {
	dbURL := os.Getenv("DB_URL")
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to connect to database: %w", err)
	}
	return db, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db, err := connectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	//stores
	adminStore := store.NewAdminStore(db)
	sessionStore := store.NewSessionStore(db)

	//middlewares
	authMiddleware := middleware.NewAuthmiddleware(adminStore, sessionStore)

	//handlers
	adminHandler := handlers.NewAdminHandler(adminStore, sessionStore)

	if err := adminStore.EnsureAdminAccountExists(); err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)

	//routes
	r.Post("/admin/auth/login", adminHandler.Login)
	r.With(authMiddleware.SessionAuth).Put("/admin/password", adminHandler.ChangePassword)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}

	log.Println("[INFO]: Server listening on ", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
