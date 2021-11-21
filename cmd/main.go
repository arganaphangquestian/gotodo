package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

var (
	APPLICATION_PORT string
	DATABASE_URL     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	APPLICATION_PORT = os.Getenv("APPLICATION_PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", APPLICATION_PORT), r)
}
