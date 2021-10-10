package db

import (
	"context"
	"github.com/Ruadgedy/simplebank/util"
	"log"
	"testing"
)

// return last transfer id
func createRandomTransfer(t *testing.T,account1,account2 Account) int64 {
	result, err := testQueries.CreateTransfer(context.Background(), CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	})
	if err != nil {
		log.Fatal(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	return lastInsertId
}

//func TestCreateTransfer(t *testing.T)  {
//	accountId1 := createRandomAccount(t)
//	accountId2 := createRandomAccount(t)
//	create
//}