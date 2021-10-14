package main

import (
	"database/sql"
	"github.com/Ruadgedy/simplebank/api"
	db "github.com/Ruadgedy/simplebank/db/sqlc"
	"github.com/Ruadgedy/simplebank/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//const (
//	dbDriver = "mysql"
//	dbSource = "root:passwd@tcp(127.0.0.1:3307)/bank?parseTime=true"
//	serverAddress = "0.0.0.0:8080"
//)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load configuration", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server,err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("can not create server:",err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}

}
