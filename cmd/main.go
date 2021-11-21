package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	todo "github.com/arganaphangquestian/gotodo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4"
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
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "OK",
		})
	})

	db, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatalln(err)
	}

	todoService := todo.TodoService{
		DB: db,
	}

	// Todo API
	todoRouter := chi.NewRouter()
	todoRouter.Get("/", todoService.GetTodos)
	todoRouter.Get("/{id}", todoService.GetTodo)
	todoRouter.Post("/", todoService.CreateTodo)
	todoRouter.Put("/{id}", todoService.UpdateTodo)
	todoRouter.Delete("/{id}", todoService.DeleteTodo)
	r.Mount("/api/todo", todoRouter)

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", APPLICATION_PORT), r)
}
