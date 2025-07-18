-- name: CreateUsageHistory :one
INSERT INTO usage_history (
    usage_history_id,
    inventaris_id,
    old_room_id,
    new_room_id,
    old_user_id,
    new_user_id,
    moved_at,
    created_at,
    created_by
) VALUES (
             @usage_history_id,
             @inventaris_id,
             @old_room_id,
             @new_room_id,
             @old_user_id,
             @new_user_id,
             @moved_at,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetUsageHistoryByID :one
SELECT *
FROM usage_history
WHERE usage_history_id = @usage_history_id
  AND deleted_at IS NULL;

-- name: ListUsageHistory :many
SELECT *
FROM usage_history
WHERE deleted_at IS NULL
ORDER BY moved_at DESC NULLS LAST, created_at DESC;

-- name: UpdateUsageHistory :one
UPDATE usage_history
SET
    inventaris_id = @inventaris_id,
    old_room_id = @old_room_id,
    new_room_id = @new_room_id,
    old_user_id = @old_user_id,
    new_user_id = @new_user_id,
    moved_at = @moved_at,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE usage_history_id = @usage_history_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteUsageHistory :exec
UPDATE usage_history
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE usage_history_id = @usage_history_id;

-- name: RestoreUsageHistory :exec
UPDATE usage_history
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE usage_history_id = @usage_history_id;
