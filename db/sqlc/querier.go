// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	// @cache-ttl 30
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	// @cache-ttl 30
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	// @cache-ttl 30
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	// @cache-ttl 30
	CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error)
	// @cache-ttl 30
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// @cache-ttl 30
	DeleteAccount(ctx context.Context, id int64) (int32, error)
	// @cache-ttl 30
	DeleteEntry(ctx context.Context, id int64) (int32, error)
	// @cache-ttl 30
	DeleteTransfer(ctx context.Context, id int64) (int32, error)
	// @cache-ttl 30
	GetAccount(ctx context.Context, id int64) (Account, error)
	// @cache-ttl 30
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	// @cache-ttl 30
	GetEntry(ctx context.Context, id int64) (Entry, error)
	// @cache-ttl 30
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	// @cache-ttl 30
	GetUser(ctx context.Context, username string) (User, error)
	// @cache-ttl 30
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	// @cache-ttl 30
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	// @cache-ttl 30
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	// @cache-ttl 30
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	// @cache-ttl 30
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error)
	// @cache-ttl 30
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error)
	// @cache-ttl 30
	UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error)
}

var _ Querier = (*Queries)(nil)
