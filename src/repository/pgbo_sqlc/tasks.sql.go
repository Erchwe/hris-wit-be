// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tasks.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
    list_id, task_order, task_name, task_type, task_priority,
    task_size, task_status, task_color, start_date, due_date,
    created_at, created_by
) VALUES (
     $1, $2, $3, $4::task_type_enum, $5::task_priority_enum,
    $6::task_size_enum, $7::task_status_enum, $8, $9, $10,
    (now() at time zone 'UTC')::TIMESTAMP, $11
) RETURNING id, task_id, list_id, task_name, task_type, task_priority, task_size, task_status, task_color, start_date, due_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, task_order
`

type CreateTaskParams struct {
	ListID       string          `json:"list_id"`
	TaskOrder    sql.NullFloat64 `json:"task_order"`
	TaskName     sql.NullString  `json:"task_name"`
	TaskType     interface{}     `json:"task_type"`
	TaskPriority interface{}     `json:"task_priority"`
	TaskSize     interface{}     `json:"task_size"`
	TaskStatus   interface{}     `json:"task_status"`
	TaskColor    sql.NullString  `json:"task_color"`
	StartDate    sql.NullTime    `json:"start_date"`
	DueDate      sql.NullTime    `json:"due_date"`
	CreatedBy    string          `json:"created_by"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ListID,
		arg.TaskOrder,
		arg.TaskName,
		arg.TaskType,
		arg.TaskPriority,
		arg.TaskSize,
		arg.TaskStatus,
		arg.TaskColor,
		arg.StartDate,
		arg.DueDate,
		arg.CreatedBy,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.TaskID,
		&i.ListID,
		&i.TaskName,
		&i.TaskType,
		&i.TaskPriority,
		&i.TaskSize,
		&i.TaskStatus,
		&i.TaskColor,
		&i.StartDate,
		&i.DueDate,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.TaskOrder,
	)
	return i, err
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, task_id, list_id, task_name, task_type, task_priority, task_size, task_status, task_color, start_date, due_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, task_order FROM tasks
WHERE task_id = $1
  AND deleted_at IS NULL
`

func (q *Queries) GetTaskByID(ctx context.Context, taskID string) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, taskID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.TaskID,
		&i.ListID,
		&i.TaskName,
		&i.TaskType,
		&i.TaskPriority,
		&i.TaskSize,
		&i.TaskStatus,
		&i.TaskColor,
		&i.StartDate,
		&i.DueDate,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.TaskOrder,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, task_id, list_id, task_name, task_type, task_priority, task_size, task_status, task_color, start_date, due_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, task_order FROM tasks
WHERE deleted_at IS NULL
ORDER BY task_order ASC
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
			&i.TaskID,
			&i.ListID,
			&i.TaskName,
			&i.TaskType,
			&i.TaskPriority,
			&i.TaskSize,
			&i.TaskStatus,
			&i.TaskColor,
			&i.StartDate,
			&i.DueDate,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.DeletedAt,
			&i.DeletedBy,
			&i.TaskOrder,
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

const restoreTask = `-- name: RestoreTask :exec
UPDATE tasks
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE task_id = $1
`

func (q *Queries) RestoreTask(ctx context.Context, taskID string) error {
	_, err := q.db.ExecContext(ctx, restoreTask, taskID)
	return err
}

const softDeleteTask = `-- name: SoftDeleteTask :exec
UPDATE tasks
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = $1
WHERE task_id = $2
`

type SoftDeleteTaskParams struct {
	DeletedBy sql.NullString `json:"deleted_by"`
	TaskID    string         `json:"task_id"`
}

func (q *Queries) SoftDeleteTask(ctx context.Context, arg SoftDeleteTaskParams) error {
	_, err := q.db.ExecContext(ctx, softDeleteTask, arg.DeletedBy, arg.TaskID)
	return err
}

const updateTask = `-- name: UpdateTask :one
UPDATE tasks
SET
    task_name = $1,
    list_id = $2,
    task_order = $3,
    task_type = $4::task_type_enum,
    task_priority = $5::task_priority_enum,
    task_size = $6::task_size_enum,
    task_status = $7::task_status_enum,
    task_color = $8,
    start_date = $9,
    due_date = $10,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $11
WHERE task_id = $12
  AND deleted_at IS NULL
RETURNING id, task_id, list_id, task_name, task_type, task_priority, task_size, task_status, task_color, start_date, due_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, task_order
`

type UpdateTaskParams struct {
	TaskName     sql.NullString  `json:"task_name"`
	ListID       string          `json:"list_id"`
	TaskOrder    sql.NullFloat64 `json:"task_order"`
	TaskType     interface{}     `json:"task_type"`
	TaskPriority interface{}     `json:"task_priority"`
	TaskSize     interface{}     `json:"task_size"`
	TaskStatus   interface{}     `json:"task_status"`
	TaskColor    sql.NullString  `json:"task_color"`
	StartDate    sql.NullTime    `json:"start_date"`
	DueDate      sql.NullTime    `json:"due_date"`
	UpdatedBy    sql.NullString  `json:"updated_by"`
	TaskID       string          `json:"task_id"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTask,
		arg.TaskName,
		arg.ListID,
		arg.TaskOrder,
		arg.TaskType,
		arg.TaskPriority,
		arg.TaskSize,
		arg.TaskStatus,
		arg.TaskColor,
		arg.StartDate,
		arg.DueDate,
		arg.UpdatedBy,
		arg.TaskID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.TaskID,
		&i.ListID,
		&i.TaskName,
		&i.TaskType,
		&i.TaskPriority,
		&i.TaskSize,
		&i.TaskStatus,
		&i.TaskColor,
		&i.StartDate,
		&i.DueDate,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.TaskOrder,
	)
	return i, err
}
