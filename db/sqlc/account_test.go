package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Ruadgedy/simplebank/util"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

// return lastInsertId
func createRandomAccount(t *testing.T) int64 {
	params := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	result, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
	if err != nil {
		log.Fatal("Create account error: ", err)
	}
	if affected, err := result.RowsAffected(); err == nil {
		fmt.Printf("Affected rows total: %d\n", affected)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastInsertId
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// creaet account
	lastInsertId := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), lastInsertId)
	if err != nil {
		log.Fatal("Get account error: ", err)
	}
	fmt.Printf("Owner:%s\n Balance:%d\n CreatedAt:%v\n Currency:%s\n", account.Owner, account.Balance, account.CreatedAt, account.Currency)
}

func TestUpdateAccount(t *testing.T) {
	lastInsertId := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), lastInsertId)
	if err != nil {
		log.Fatal("Get account failed: ", err)
	}
	testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		Balance: util.RandomMoney(),
		ID:      lastInsertId,
	})

	new_account, err := testQueries.GetAccount(context.Background(), lastInsertId)
	if err != nil {
		log.Fatal("Get account failed: ", err)
	}
	fmt.Printf("The Origin balance: %d, the newly balance: %d\n", account.Balance, new_account.Balance)
}

func TestDeleteAccount(t *testing.T) {
	lastInsertId := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), lastInsertId)
	if err != nil {
		log.Fatal("Get account error: ", err)
	}
	fmt.Printf("Insert new account: %v\n", account)

	fmt.Println("Begin delete Account where id =", account.ID)

	testQueries.DeleteAccount(context.Background(), account.ID)

	account, err = testQueries.GetAccount(context.Background(), lastInsertId)
	if err != nil && err == sql.ErrNoRows {
		fmt.Println("Delete Account Success")
	}
}

func TestListAccounts(t *testing.T) {
	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No such rows:", err)
			return
		}
		log.Fatal(err)
	}
	for _, account := range accounts {
		fmt.Println(account)
	}
}
