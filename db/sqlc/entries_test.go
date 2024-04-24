package db

import (
	"context"
	"testing"

	"github.com/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestEntries(t *testing.T) {
	t.Parallel()
	testCreateEntry(t)
	testGetEntry(t)
	testListEntries(t)
	testUpdateEntry(t)
	testDeleteEntry(t)
}

func testCreateEntry(t *testing.T) Entry {
	args := CreateEntryParams{
		AccountID: utils.RandomINT(1, 4),
		Amount:    utils.RandomINT(1, 200),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.Amount, entry.Amount)
	require.Equal(t, args.AccountID, entry.AccountID)

	require.NotZero(t, entry.CreateAt)
	require.NotZero(t, entry.ID)

	return entry
}

func testGetEntry(t *testing.T) Entry {

	entry, err := testQueries.GetEntry(context.Background(), utils.RandomINT(1, 4))
	if err != nil {
		require.Empty(t, entry)
	}

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.CreateAt)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.Amount)

	return entry

}

func testListEntries(t *testing.T) []Entry {
	args := ListEntriesParams{
		Limit:  1,
		Offset: 1,
	}
	entries, err := testQueries.ListEntries(context.Background(), args)
	if err != nil {
		require.Empty(t, entries)
	}
	require.NoError(t, err)

	return entries

}

func testUpdateEntry(t *testing.T) {
	args := UpdateEntryParams{
		ID:     utils.RandomINT(1, 2),
		Amount: utils.RandomINT(1, 200),
	}

	entry, err := testQueries.UpdateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.Amount, entry.Amount)
	require.Equal(t, args.ID, entry.ID)

	require.NotZero(t, entry.CreateAt)
	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.ID)
}

func testDeleteEntry(t *testing.T) {

	entry, err := testQueries.DeleteEntry(context.Background(), utils.RandomINT(1, 2))

	require.NoError(t, err)
	require.Zero(t, entry)
}
