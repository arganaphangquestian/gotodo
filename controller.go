package todo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/arganaphangquestian/gotodo/data"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *string     `json:"error,omitempty"`
}

type TodoService struct {
	DB *pgx.Conn
}

func (s *TodoService) GetTodos(w http.ResponseWriter, r *http.Request) {
	queries := data.New(s.DB)
	todos, _ := queries.ListTodos(context.Background())
	json.NewEncoder(w).Encode(response{
		Message: "Get Todos",
		Data: map[string]interface{}{
			"todos": todos,
		},
	})
}

func (s *TodoService) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Message: "Failed Get Todo",
		})
		return
	}
	queries := data.New(s.DB)
	todos, err := queries.GetTodo(context.Background(), int64(id))
	json.NewEncoder(w).Encode(response{
		Message: "Get Todo",
		Data: map[string]interface{}{
			"todos": todos,
		},
	})
}

type TodoRequest struct {
	*data.Todo
}

func (a *TodoRequest) Bind(r *http.Request) error {
	if a.Todo == nil {
		return errors.New("missing required Todo fields")
	}
	return nil
}

func (s *TodoService) CreateTodo(w http.ResponseWriter, r *http.Request) {
	fields := &TodoRequest{}
	if err := render.Bind(r, fields); err != nil {
		json.NewEncoder(w).Encode(response{
			Message: "Failed Create Todo",
		})
		return
	}
	queries := data.New(s.DB)
	todo, _ := queries.CreateTodo(context.Background(), data.CreateTodoParams{
		Title: fields.Title,
		Done:  fields.Done,
	})
	json.NewEncoder(w).Encode(response{
		Message: "Create Todo",
		Data: map[string]interface{}{
			"todo": todo,
		},
	})
}

func (s *TodoService) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Message: "Failed Update Todo",
		})
		return
	}
	fields := &TodoRequest{}
	if err := render.Bind(r, fields); err != nil {
		json.NewEncoder(w).Encode(response{
			Message: "Failed Update Todo",
		})
		return
	}
	queries := data.New(s.DB)
	todo, err := queries.UpdateTodo(context.Background(), data.UpdateTodoParams{
		ID:    int64(id),
		Title: fields.Title,
		Done:  fields.Done,
	})
	json.NewEncoder(w).Encode(response{
		Message: "Update Todo",
		Data: map[string]interface{}{
			"todo": todo,
		},
	})
}

func (s *TodoService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Message: "Failed Delete Todo",
		})
		return
	}
	queries := data.New(s.DB)
	err = queries.DeleteTodo(context.Background(), int64(id))
	json.NewEncoder(w).Encode(response{
		Message: "Delete Todo",
	})
}
