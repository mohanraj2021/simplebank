-- name: CreateEntry :one
-- @cache-ttl 30
INSERT INTO entries (
  account_id,amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
-- @cache-ttl 30
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
-- @cache-ttl 30
SELECT * FROM entries
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
-- @cache-ttl 30
UPDATE entries SET 
amount = $2
WHERE id = $1
RETURNING *;


-- name: DeleteEntry :one
-- @cache-ttl 30
DELETE FROM entries WHERE id = $1
RETURNING 0;