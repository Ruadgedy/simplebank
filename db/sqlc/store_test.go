package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferTx(t *testing.T)  {
	store := NewStore(testDB)

	account1Id := createRandomAccount(t)
	account2Id := createRandomAccount(t)

	account1, _ := testQueries.GetAccount(context.Background(), account1Id)
	account2, _ := testQueries.GetAccount(context.Background(), account2Id)

	// run n concurrent transfer transaction
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: account1.ID,
				ToAccountId:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i:=0; i < n; i++ {
		err := <-errs
		require.NoError(t,err)

		result := <- results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t,transfer)

		fmt.Printf("%v\n", result)
	}
}
