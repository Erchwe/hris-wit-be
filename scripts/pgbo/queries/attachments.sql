-- name: CreateAttachment :one
INSERT INTO attachments (
    attach_id, task_id, attach_name, attach_url,
    created_at, created_by
) VALUES (
    @attach_id, @task_id, @attach_name, @attach_url,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetAttachmentByID :one
SELECT * FROM attachments
WHERE attach_id = @attach_id
  AND deleted_at IS NULL;

-- name: ListAttachments :many
SELECT * FROM attachments
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateAttachment :one
UPDATE attachments
SET
    task_id = @task_id,
    attach_name = @attach_name,
    attach_url = @attach_url,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE attach_id = @attach_id
  AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteAttachment :exec
UPDATE attachments
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE attach_id = @attach_id;

-- name: RestoreAttachment :exec
UPDATE attachments
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE attach_id = @attach_id;

