// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package data

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (
  title, done
) VALUES (
  $1, $2
)
RETURNING id, title, done
`

type CreateTodoParams struct {
	Title string
	Done  sql.NullBool
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo, arg.Title, arg.Done)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Done)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, done FROM todos
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodo, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Done)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, done FROM todos
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Title, &i.Done); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos SET title = $1, done = $2 WHERE id = $3
RETURNING id, title, done
`

type UpdateTodoParams struct {
	Title string
	Done  sql.NullBool
	ID    int64
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodo, arg.Title, arg.Done, arg.ID)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Done)
	return i, err
}