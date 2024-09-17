-- name: CreateTask :one
INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks WHERE id = $1;

-- name: ListTasks :many
SELECT * FROM tasks;

-- name: UpdateTask :one
UPDATE tasks SET title = $2, description = $3, status = $4 WHERE id = $1 RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: CreateAuthor :one
INSERT INTO authors (name) VALUES ($1) RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM authors;

-- name: AssignAuthorToTask :exec
INSERT INTO task_authors (task_id, author_id) VALUES ($1, $2);
