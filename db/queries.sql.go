// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
)

const assignAuthorToTask = `-- name: AssignAuthorToTask :exec
INSERT INTO task_authors (task_id, author_id) VALUES ($1, $2)
`

type AssignAuthorToTaskParams struct {
	TaskID   int32 `json:"task_id"`
	AuthorID int32 `json:"author_id"`
}

func (q *Queries) AssignAuthorToTask(ctx context.Context, arg AssignAuthorToTaskParams) error {
	_, err := q.db.ExecContext(ctx, assignAuthorToTask, arg.TaskID, arg.AuthorID)
	return err
}

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, name)
	var i Author
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, title, description, status
`

type CreateTaskParams struct {
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.Title, arg.Description, arg.Status)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name FROM authors WHERE id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, id int32) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getTask = `-- name: GetTask :one
SELECT id, title, description, status FROM tasks WHERE id = $1
`

func (q *Queries) GetTask(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name FROM authors
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTasks = `-- name: ListTasks :many
SELECT id, title, description, status FROM tasks
`

func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :one
UPDATE tasks SET title = $2, description = $3, status = $4 WHERE id = $1 RETURNING id, title, description, status
`

type UpdateTaskParams struct {
	ID          int32          `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTask,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Status,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
	)
	return i, err
}
