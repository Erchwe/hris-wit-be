-- name: CreateList :one
INSERT INTO lists (
    project_id, list_name, list_order,
    created_at, created_by
) VALUES (
     @project_id, @list_name, @list_order,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING list_id;

-- name: GetListByID :one
SELECT * FROM lists
WHERE list_id = @list_id
  AND deleted_at IS NULL;

-- name: GetListByProjectIDAndListID :one
SELECT * FROM lists
WHERE list_id = @list_id
AND project_id = @project_id
AND deleted_at IS NULL;

-- name: ListListsByProjectID :many
SELECT * FROM lists
where project_id = @project_id
  AND deleted_at IS NULL;

-- name: ListLists :many
SELECT * FROM lists
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateList :one
UPDATE lists
SET
    list_order = @list_order,
    list_name = @list_name,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE list_id = @list_id
  AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteList :exec
UPDATE lists
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE list_id = @list_id;

-- name: RestoreList :exec
UPDATE lists
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE list_id = @list_id;
