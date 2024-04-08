package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	result := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			transferTxResult, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			result <- transferTxResult
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		res := <-result
		require.NoError(t, err)

		require.NotEmpty(t, res)

		require.NotEmpty(t, res.Transfer)

		require.Equal(t, account1.ID, res.Transfer.FromAccountID)
		require.Equal(t, account2.ID, res.Transfer.ToAccountID)
		require.Equal(t, amount, res.Transfer.Amount)
		require.NotZero(t, res.Transfer.ID)
		require.NotZero(t, res.Transfer.CreateAt)

		_, gterr := store.GetTransfer(context.Background(), res.Transfer.ID)

		require.NoError(t, gterr)

		fromEntry := res.FromEntry

		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.CreateAt)
		require.NotZero(t, fromEntry.ID)

		toEntry := res.ToEntry

		require.NotEmpty(t, toEntry)
		require.Equal(t, account1.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.CreateAt)
		require.NotZero(t, toEntry.ID)
	}
}
