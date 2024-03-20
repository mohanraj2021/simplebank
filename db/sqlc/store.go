package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func (store *Store) exeTrx(ctx context.Context, fn func(q *Queries) error) error {

	tx, terr := store.db.BeginTx(ctx, nil)

	if terr != nil {
		return terr
	}
	qrs := New(tx)
	if qerr := fn(qrs); qerr != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("qery error %v, rollback error %v", qerr, rberr)
		}
		return qerr
	}
	return tx.Commit()

}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	store.exeTrx(ctx, func(q *Queries) error {
		var terr error
		result.Transfer, terr = q.CreateTransfers(ctx, CreateTransfersParams(args))
		if terr != nil {
			return terr
		}

		result.FromEntry, terr = q.CreateEntry(ctx, CreateEntryParams{Amount: -args.Amount, AccountID: args.FromAccountID})
		if terr != nil {
			return terr
		}

		result.ToEntry, terr = q.CreateEntry(ctx, CreateEntryParams{Amount: args.Amount, AccountID: args.ToAccountID})
		if terr != nil {
			return terr
		}

		return nil
	})

	return result, nil
}
