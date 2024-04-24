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

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
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

var txKey = struct{}{}

func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	var txName = ctx.Value(txKey)
	store.exeTrx(ctx, func(q *Queries) error {
		var terr error

		fmt.Println(txName, "Create Transfer")
		result.Transfer, terr = q.CreateTransfers(ctx, CreateTransfersParams(args))
		if terr != nil {
			return terr
		}

		fmt.Println(txName, "Create From entry")
		result.FromEntry, terr = q.CreateEntry(ctx, CreateEntryParams{Amount: -args.Amount, AccountID: args.FromAccountID})
		if terr != nil {
			return terr
		}

		fmt.Println(txName, "Create TO entry")
		result.ToEntry, terr = q.CreateEntry(ctx, CreateEntryParams{Amount: args.Amount, AccountID: args.ToAccountID})
		if terr != nil {
			return terr
		}

		fmt.Println(txName, "from getaccount for update")
		account1, err := q.GetAccountForUpdate(context.Background(), args.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "to getaccount for update")
		result.FromAccount, err = q.UpdateAccount(context.Background(), UpdateAccountParams{
			ID:      args.FromAccountID,
			Balance: account1.Balance - args.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "from account update")
		account2, err := q.GetAccountForUpdate(context.Background(), args.ToAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "to account update")
		result.ToAccount, err = q.UpdateAccount(context.Background(), UpdateAccountParams{
			ID:      args.ToAccountID,
			Balance: account2.Balance + args.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, nil
}
