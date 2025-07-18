-- name: CreateClient :one
INSERT INTO client (
    client_name, shipment_address, billing_address,
    created_at, created_by
) VALUES (
     @client_name, @shipment_address, @billing_address,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING client_id;

-- name: GetClientByID :one
SELECT * FROM client
WHERE client_id = @client_id
  AND deleted_at IS NULL;

-- name: ListClients :many
SELECT * FROM client
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateClient :one
UPDATE client
SET
    client_name = @client_name,
    shipment_address = @shipment_address,
    billing_address = @billing_address,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE client_id = @client_id
  AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteClient :exec
UPDATE client
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE client_id = @client_id
  AND deleted_at IS NULL;

-- name: RestoreClient :exec
UPDATE client
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE client_id = @client_id;
