package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gitnoober/No-Bank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    utils.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)

	require.NotZero(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	const limit int = 5
	const offset int = 5
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, limit)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)

}
