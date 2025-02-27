package db

import (
	"context"
	"testing"

	"github.com/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomINT(3, 6),
		Currency: utils.RandomCurrency(),
	}

	acc, aerr := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, aerr)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreateAt)

	return acc

}

func TestAccount(t *testing.T) {
	createRandomAccount(t)
	upDateAccount(t)
	getRandomAccount(t)
	deleteRandomAccount(t)
	getRandAccountForUpdate(t)
	getListAccount(t)
}

func upDateAccount(t *testing.T) Account {
	args := UpdateAccountParams{
		ID:      1,
		Balance: 125,
	}
	acc, aerr := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, aerr)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.ID, args.ID)
	require.Equal(t, acc.Balance, args.Balance)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreateAt)

	return acc
}

func getRandomAccount(t *testing.T) Account {
	accountId := utils.RandomINT(1, 4)
	acc, aerr := testQueries.GetAccount(context.Background(), int64(accountId))

	require.NoError(t, aerr)
	require.NotEmpty(t, acc)

	require.NotZero(t, accountId, acc.ID)
	require.NotEmpty(t, acc.Owner)
	require.NotEmpty(t, acc.Currency)

	return acc
}

func deleteRandomAccount(t *testing.T) {
	accountId := utils.RandomINT(1, 4)
	acc, aerr := testQueries.DeleteAccount(context.Background(), int64(accountId))

	require.NoError(t, aerr)
	// require.NotZero(t, acc)

	require.NotEmpty(t, acc)
}

func getRandAccountForUpdate(t *testing.T) {
	accountId := utils.RandomINT(1, 4)
	acc, aerr := testQueries.GetAccountForUpdate(context.Background(), accountId)
	require.NoError(t, aerr)

	require.NotEmpty(t, acc)
}

func getListAccount(t *testing.T) {
	args := ListAccountsParams{
		Limit:  3,
		Offset: 0,
	}
	acc, aerr := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, aerr)

	require.NotEmpty(t, acc)
	require.NotEqual(t, 0, len(acc))
}
