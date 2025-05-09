// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
fullname,
email,
hashedpassword
) VALUES (
  $1, $2,$3,$4
)
RETURNING username, fullname, email, hashedpassword, created_at, updated_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Hashedpassword string `json:"hashedpassword"`
}

// @cache-ttl 30
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Fullname,
		arg.Email,
		arg.Hashedpassword,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Fullname,
		&i.Email,
		&i.Hashedpassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, fullname, email, hashedpassword, created_at, updated_at FROM users
WHERE username = $1 LIMIT 1
`

// @cache-ttl 30
func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Fullname,
		&i.Email,
		&i.Hashedpassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE users SET 
hashedpassword = $2,
updated_at = Now()
WHERE username = $1
RETURNING username, fullname, email, hashedpassword, created_at, updated_at
`

type UpdatePasswordParams struct {
	Username       string `json:"username"`
	Hashedpassword string `json:"hashedpassword"`
}

// @cache-ttl 30
func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.Username, arg.Hashedpassword)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Fullname,
		&i.Email,
		&i.Hashedpassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
