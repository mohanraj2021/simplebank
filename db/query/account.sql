-- name: CreateAccount :one
-- @cache-ttl 30
INSERT INTO account (
  owner,balance,curreny
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetAccount :one
-- @cache-ttl 30
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
-- @cache-ttl 30
SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: ListAccounts :many
-- @cache-ttl 30
SELECT * FROM account
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
-- @cache-ttl 30
UPDATE account SET 
balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :one
-- @cache-ttl 30
DELETE FROM account WHERE id = $1
RETURNING 0;