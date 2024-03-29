package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	// db "github.com/gitnoober/No-Bank/db/sqlc"
	"github.com/gitnoober/No-Bank/utils"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var testQueries *Queries
var testDB *sql.DB

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	config, config_err := utils.LoadConfig("../..")
	if config_err != nil {
		log.Fatal("cannot load config:", config_err)
	}
	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run()) // running testing object and os.exit would report it back to the runner
}

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err) // assert if error is nil
	require.NotEmpty(t, account)

	// check if result's params are similar to args object
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func createRandomEntry(t *testing.T) Entry {
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
	return entry
}
