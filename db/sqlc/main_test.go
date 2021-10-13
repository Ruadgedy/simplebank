package db

import (
	"database/sql"
	"github.com/Ruadgedy/simplebank/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
	"time"
)

var testQueries *Queries
var testDB *sql.DB

//const (
//	dbDriver = "mysql"
//	dbSource = "root:passwd@tcp(127.0.0.1:3307)/bank?parseTime=true"
//)

func TestMain(m *testing.M) {
	// Open函数需要注意：Driver这里是空实现，没有具体的驱动。如果要使用mysql驱动，需要添加上面的依赖
	var err error
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("can not load config ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to database: ", err)
	}
	testDB.SetConnMaxLifetime(time.Minute * 10)
	defer testDB.Close()
	if err = testDB.Ping(); err != nil {
		log.Fatalln("Open database connection failed: ", err)
	}
	log.Println("Start testing....")
	testQueries = New(testDB)
	os.Exit(m.Run())
}
