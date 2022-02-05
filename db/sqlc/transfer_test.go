package db

import (
	"context"
	"testing"
	"time"

	"github.com/oaraujocesar/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, from_account.ID, transfer.FromAccountID)
	require.Equal(t, to_account.ID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)

	transfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, createdTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, createdTransfer.Amount, transfer.Amount)
	require.Equal(t, createdTransfer.ID, transfer.ID)

	require.WithinDuration(t, createdTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}
