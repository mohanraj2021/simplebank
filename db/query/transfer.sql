-- name: CreateTransfers :one
-- @cache-ttl 30
INSERT INTO transfers (
  from_account_id,to_account_id,amount
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetTransfer :one
-- @cache-ttl 30
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
-- @cache-ttl 30
SELECT * FROM transfers
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
-- @cache-ttl 30
UPDATE transfers SET 
amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :one
-- @cache-ttl 30
DELETE FROM transfers WHERE id = $1
RETURNING 0;