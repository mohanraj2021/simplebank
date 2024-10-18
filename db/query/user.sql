-- name: CreateUser :one
-- @cache-ttl 30
INSERT INTO users (
  username,
fullname,
email,
hashedpassword
) VALUES (
  $1, $2,$3,$4
)
RETURNING *;


-- name: GetUser :one
-- @cache-ttl 30
SELECT * FROM users
WHERE username = $1 LIMIT 1;


-- name: UpdatePassword :one
-- @cache-ttl 30
UPDATE users SET 
hashedpassword = $2,
updated_at = Now()
WHERE username = $1
RETURNING *;