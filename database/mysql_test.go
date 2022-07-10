package database_test

import (
	"fmt"
	"log"
	"testing"
	"with_go/database"
)

func TestGetServerConfig(t *testing.T) {
	config, _ := database.GetServerConfig("mysql_Oracle_ex")
	if config == nil {
		t.Error("Wrong result")
	}
}

func TestMysqlConn(t *testing.T) {
	config, err := database.GetServerConfig("mysql_Oracle_ex")
	if err != nil {
		log.Fatal(err)
	}
	dbName := "coin"
	// fmt.Printf("%s:%s@tcp(%s:%d)/%s", config["user"], config["passwd"], config["host"], config["port"], dbName)
	dsn, _ := database.ConnectDb(config, dbName)
	conn, err := database.OpenDb(dsn)
	if conn == nil {
		t.Error("Wrong result")
	}
}

func TestMysqlFind(t *testing.T) {
	config, err := database.GetServerConfig("mysql_Oracle_ex")
	if err != nil {
		log.Fatal(err)
	}
	dbName := "coin"
	// fmt.Printf("%s:%s@tcp(%s:%d)/%s", config["user"], config["passwd"], config["host"], config["port"], dbName)
	dsn, dbms := database.ConnectDb(config, dbName)
	conn, err := database.OpenDb(dsn)
	if conn == nil {
		t.Error("Wrong result")
	}
	fmt.Println(database.Find([]string{"eng", "kor", "ticker"}, "coins", conn, dbms))

}

// func TestMysqlInsert(t *testing.T) {
// 	config, err := database.GetServerConfig("mysql_Oracle_ex")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dbName := "coin"
// 	// fmt.Printf("%s:%s@tcp(%s:%d)/%s", config["user"], config["passwd"], config["host"], config["port"], dbName)
// 	dsn, dbms := database.ConnectDb(config, dbName)
// 	conn, err := database.OpenDb(dsn)
// 	if conn == nil {
// 		t.Error("Wrong result")
// 	}

// 	// INSERT
// 	data := []map[string]interface{}{
// 		map[string]interface{}{"eng": "Ripple", "kor": "리플", "ticker": "KRW-XRP"},
// 		map[string]interface{}{"eng": "Terra", "kor": "테라", "ticker": "KRW-TER"},
// 	}
// 	fmt.Println(setInsertKeyValsFromMaps(data))
// 	database.Insert(data, "coins", conn, dbms)
// }
