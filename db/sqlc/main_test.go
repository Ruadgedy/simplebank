package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB
const (
	dbDriver = "mysql"
	dbSource = "root:passwd@tcp(127.0.0.1:3307)/bank?parseTime=true"
)

func TestMain(m *testing.M) {
	// Open函数需要注意：Driver这里是空实现，没有具体的驱动。如果要使用mysql驱动，需要添加上面的依赖
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to database: ",err)
	}
	defer testDB.Close()
	if err = testDB.Ping(); err!=nil {
		log.Fatalln("Open database connection failed: ", err)
	}
	log.Println("Start testing....")
	testQueries = New(testDB)
	os.Exit(m.Run())
}
