package db

import (
	"context"
	"testing"
	"time"

	"github.com/oaraujocesar/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	createRandomTransfer(t, from_account, to_account)
}

func TestGetTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	createdTransfer := createRandomTransfer(t, from_account, to_account)

	transfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, createdTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, createdTransfer.Amount, transfer.Amount)
	require.Equal(t, createdTransfer.ID, transfer.ID)

	require.WithinDuration(t, createdTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {

	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from_account, to_account)
	}

	arg := ListTransfersParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == from_account.ID || transfer.ToAccountID == to_account.ID)
	}
}
