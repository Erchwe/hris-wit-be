-- name: CreateTask :one
INSERT INTO tasks (
    list_id, task_order, task_name, task_type, task_priority,
    task_size, task_status, task_color, start_date, due_date,
    created_at, created_by
) VALUES (
     @list_id, @task_order, @task_name, @task_type::task_type_enum, @task_priority::task_priority_enum,
    @task_size::task_size_enum, @task_status::task_status_enum, @task_color, @start_date, @due_date,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE task_id = @task_id
  AND deleted_at IS NULL;

-- name: ListTasks :many
SELECT * FROM tasks
WHERE deleted_at IS NULL
ORDER BY task_order ASC;

-- name: UpdateTask :one
UPDATE tasks
SET
    task_name = @task_name,
    list_id = @list_id,
    task_order = @task_order,
    task_type = @task_type::task_type_enum,
    task_priority = @task_priority::task_priority_enum,
    task_size = @task_size::task_size_enum,
    task_status = @task_status::task_status_enum,
    task_color = @task_color,
    start_date = @start_date,
    due_date = @due_date,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE task_id = @task_id
  AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteTask :exec
UPDATE tasks
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE task_id = @task_id;

-- name: RestoreTask :exec
UPDATE tasks
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE task_id = @task_id;
