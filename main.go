package main

import (
	"database/sql"
	"github.com/Ruadgedy/simplebank/api"
	db "github.com/Ruadgedy/simplebank/db/sqlc"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	dbDriver = "mysql"
	dbSource = "root:passwd@tcp(127.0.0.1:3307)/bank?parseTime=true"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}

}
