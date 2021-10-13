package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides a function to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx execute a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("tx error: %v", err)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

//TransferResult is the result of the transfer transaction.
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct {
}{}

// TransferTx performs a money transfer from one account to the other.
// It creates a transfer record,and account entries,and update accounts' balance within a single database transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		sql_result, err := q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		lastInsertId, err := sql_result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		transfer, _ := q.GetTransfer(ctx, lastInsertId)
		result.Transfer = transfer

		sql_result, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})
		if err != nil {
			log.Fatal(err)
		}
		lastInsertId, _ = sql_result.LastInsertId()
		entry, _ := q.GetEntry(ctx, lastInsertId)
		result.FromEntry = entry

		sql_result, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			log.Fatal(err)
		}
		lastInsertId, _ = sql_result.LastInsertId()
		entry, _ = q.GetEntry(ctx, lastInsertId)
		result.ToEntry = entry

		// TODO:update accounts' balance
		account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountId)
		if err != nil {
			return err
		}
		err = q.UpdateAccount(ctx, UpdateAccountParams{
			Balance: account1.Balance - arg.Amount,
			ID:      arg.FromAccountId,
		})
		if err != nil {
			return err
		}

		account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountId)
		if err != nil {
			return err
		}
		err = q.UpdateAccount(ctx, UpdateAccountParams{
			Balance: account2.Balance + arg.Amount,
			ID:      arg.ToAccountId,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
