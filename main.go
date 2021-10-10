package main

import (
	"database/sql"
	"github.com/Ruadgedy/simplebank/api"
	db "github.com/Ruadgedy/simplebank/db/sqlc"
	"log"
)

const (
	dbDriver = "mysql"
	dbSource = "root:passwd@tcp(127.0.0.1:3307)/bank?parseTime=true"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	db.NewStore(conn)

}
