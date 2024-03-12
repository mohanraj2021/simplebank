package db

import (
	"context"
	"testing"

	"github.com/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	transferRandomAccount(t)
	// getRandomTransfer(t)
	// deleteRandomTransfer(t)
}

func transferRandomAccount(t *testing.T) Transfer {
	args := CreateTransfersParams{
		Amount:        utils.RandomINT(1, 300),
		FromAccountID: utils.RandomINT(1, 4),
		ToAccountID:   utils.RandomINT(1, 4),
	}

	transfer, terr := testQueries.CreateTransfers(context.Background(), args)
	if terr != nil {
		require.Error(t, terr)
		require.Empty(t, transfer)
		return transfer
	}

	require.NoError(t, terr)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.Amount, transfer.Amount)
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreateAt)

	return transfer

}

func getRandomTransfer(t *testing.T) Transfer {
	transferId := utils.RandomINT(1, 2)

	transfer, terr := testQueries.GetTransfer(context.Background(), transferId)

	require.NoError(t, terr)
	require.NotEmpty(t, transfer)

	require.NotZero(t, transfer.FromAccountID)
	require.NotZero(t, transfer.ToAccountID)
	return transfer
}

func deleteRandomTransfer(t *testing.T) {
	transferID := utils.RandomINT(1, 2)
	transfer, terr := testQueries.DeleteTransfer(context.Background(), transferID)

	require.NoError(t, terr)
	require.Zero(t, transfer)
}
