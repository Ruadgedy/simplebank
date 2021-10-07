-- name: CreateAccount :execresult
insert into accounts(
    owner, balance, currency
) values (
    ?,?,?
);

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1
FOR UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
limit ?
offset ?;

-- name: UpdateAccount :exec
UPDATE accounts SET balance = ?
WHERE id = ?;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = ?;