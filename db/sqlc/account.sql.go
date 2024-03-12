// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
  owner,balance,curreny
) VALUES (
  $1, $2,$3
)
RETURNING id, owner, balance, curreny, create_at
`

type CreateAccountParams struct {
	Owner   string `json:"owner"`
	Balance int64  `json:"balance"`
	Curreny string `json:"curreny"`
}

// @cache-ttl 30
func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Curreny)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Curreny,
		&i.CreateAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :one
DELETE FROM account WHERE id = $1
RETURNING 0
`

// @cache-ttl 30
func (q *Queries) DeleteAccount(ctx context.Context, id int64) (int32, error) {
	row := q.db.QueryRowContext(ctx, deleteAccount, id)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, curreny, create_at FROM account
WHERE id = $1 LIMIT 1
`

// @cache-ttl 30
func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Curreny,
		&i.CreateAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, balance, curreny, create_at FROM account
ORDER BY id 
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// @cache-ttl 30
func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Curreny,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account SET 
balance = $2
WHERE id = $1
RETURNING id, owner, balance, curreny, create_at
`

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

// @cache-ttl 30
func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Curreny,
		&i.CreateAt,
	)
	return i, err
}
