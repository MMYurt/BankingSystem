package db

import (
	"BankingSystem/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, acc1 Account, acc2 Account) Transfer {
	amount := util.RandomMoney()

	arg := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        amount,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, acc1.ID)
	require.Equal(t, transfer.ToAccountID, acc2.ID)
	require.Equal(t, transfer.Amount, amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	createRandomTransfer(t, acc1, acc2)

}

func TestGetTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, acc1, acc2)

	getTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)

	require.Equal(t, transfer.Amount, getTransfer.Amount)
	require.Equal(t, transfer.FromAccountID, getTransfer.FromAccountID)
	require.Equal(t, transfer.ID, getTransfer.ID)
	require.Equal(t, transfer.ToAccountID, getTransfer.ToAccountID)

	require.WithinDuration(t, transfer.CreatedAt, getTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, acc1, acc2)
	}

	arg := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Limit:         5,
		Offset:        5,
	}
	transferList, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transferList)

	require.Len(t, transferList, 5)

	for _, transfer := range transferList {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, acc1.ID)
		require.Equal(t, transfer.ToAccountID, acc2.ID)
	}
}
