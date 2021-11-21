-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos;

-- name: CreateTodo :one
INSERT INTO todos (
  title, done
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos SET title = $1, done = $2 WHERE id = $3
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;