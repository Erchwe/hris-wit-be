-- name: CreateProject :one
INSERT INTO projects (
    client_id, project_name, project_desc, project_status,
    project_priority, project_color, start_date, due_date,
    created_at, created_by
) VALUES (
     @client_id, @project_name, @project_desc, @project_status::project_status_enum,
    @project_priority, @project_color, @start_date, @due_date,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING project_id;

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE project_id = @project_id
  AND deleted_at IS NULL;

-- name: ListProjects :many
SELECT * FROM projects
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateProject :one
UPDATE projects
SET
    project_name = @project_name,
    project_desc = @project_desc,
    project_status = @project_status::project_status_enum,
    project_priority = @project_priority,
    project_color = @project_color,
    start_date = @start_date,
    due_date = @due_date,
    client_id = @client_id,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE project_id = @project_id
  AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteProject :exec
UPDATE projects
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE project_id = @project_id;

-- name: RestoreProject :exec
UPDATE projects
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE project_id = @project_id;
