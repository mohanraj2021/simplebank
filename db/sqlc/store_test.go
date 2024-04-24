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

	n := 8
	amount := int64(10)

	errs := make(chan error)
	result := make(chan TransferTxResult)
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		// txValue := fmt.Sprintf("tx %d", i+1)
		// ctx := context.WithValue(context.Background(), txKey, txValue)
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
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.CreateAt)
		require.NotZero(t, toEntry.ID)

		fromAccount := res.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, account1.ID)

		toAccount := res.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, account2.ID)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAccount1, err1 := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err1)

	updatedAccount2, err2 := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err2)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDb)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 50
	amount := int64(10)

	errs := make(chan error)
	for i := 0; i < n; i++ {
		// txValue := fmt.Sprintf("tx %d", i+1)
		// ctx := context.WithValue(context.Background(), txKey, txValue)
		fromAccountId := account1.ID
		toAccountId := account2.ID

		if i%2 == 1 {
			fromAccountId = account2.ID
			toAccountId = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountId,
				ToAccountID:   toAccountId,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	updatedAccount1, err1 := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err1)

	updatedAccount2, err2 := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err2)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}
