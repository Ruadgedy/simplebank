package db

import (
	"context"
	"example.com/simplebank/util"
	"fmt"
	"log"
	"testing"
)

// return last inserted entry id
func createRandomEntry(t *testing.T,account Account) int64{
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	result, err := testQueries.CreateEntry(context.Background(), arg)
	if err != nil {
		log.Fatal("Create Entry failed: ", err)
	}
	lastInsertId, _ := result.LastInsertId()
	fmt.Println("LastInsertId: ",lastInsertId)
	return lastInsertId
}

func TestCreateEntry(t *testing.T)  {
	lastInsertId := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), lastInsertId)
	if err != nil {
		log.Fatalln(err)
	}
	createRandomEntry(t,account)
}

func TestGetEntry(t *testing.T)  {
	lastInsertId := createRandomAccount(t)
	account, _ := testQueries.GetAccount(context.Background(), lastInsertId)
	lastEntryId := createRandomEntry(t, account)
	entry, err := testQueries.GetEntry(context.Background(), lastEntryId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(entry)
}

func TestListEntries(t *testing.T)  {
	lastInsertId := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), lastInsertId)
	if err != nil {
		log.Fatalln(err)
	}
	_ = createRandomEntry(t, account)
	entries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    0,
	})

	for _, entry := range entries {
		fmt.Println(entry)
	}
}
