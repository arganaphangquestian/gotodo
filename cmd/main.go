package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	todo "github.com/arganaphangquestian/gotodo"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "OK",
		})
	}).Methods("GET")

	db, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatalln(err)
	}

	todoService := todo.TodoService{
		DB: db,
	}

	r.HandleFunc("/api/todo", todoService.GetTodos).Methods("Get")
	r.HandleFunc("/api/todo/{id}", todoService.GetTodo).Methods("Get")
	r.HandleFunc("/api/todo", todoService.CreateTodo).Methods("Post")
	r.HandleFunc("/api/todo/{id}", todoService.UpdateTodo).Methods("Put")
	r.HandleFunc("/api/todo/{id}", todoService.DeleteTodo).Methods("Delete")

	fmt.Printf("Server running at http://127.0.0.1:%s\n", APPLICATION_PORT)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", APPLICATION_PORT), r)
}
