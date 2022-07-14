// https://1minute-before6pm.tistory.com/19

package database

import (
	"database/sql"
	"fmt"

	// "strings"
	// "time"

	_ "github.com/go-sql-driver/mysql"

	// "with_go/env"
	"with_go/env"
)

const (
	YAML_NAME = "database_conn"
)

//DB접속정보를 가지고 있는 객체를 정의합니다.
type DBMS struct {
	User   string `yaml:"user"`
	Passwd string `yaml:"password"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"` // TODO: string -> int
	// db      string // TODO: string -> int
	// Charset string `yaml:"charset"`
}

func connectDb(config map[string]interface{}, dbName string) (*sql.DB, *DBMS, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config["user"], config["password"], config["host"], config["port"], dbName)

	dbms := &DBMS{User: config["user"].(string), Passwd: config["password"].(string), Host: config["host"].(string), Port: config["port"].(int)}
	conn, err := sql.Open("mysql", dsn)
	return conn, dbms, err
}

func ConnectDb(serverName, dbName string) (*sql.DB, *DBMS, error) {
	config := env.GetConfigYaml(YAML_NAME, serverName)
	conn, dbms, err := connectDb(config, dbName)
	return conn, dbms, err
}

// DBMS method
//쿼리하고 결과를 반환합니다.
func (dbms DBMS) MySQLExecQuery(query string, db *sql.DB) ([]map[string]interface{}, error) {
	defer db.Close()

	//db를 통해 sql문을 실행 시킨다.
	rows, err := db.Query(query)
	// 함수가 종료되면 rows도 Close한다.
	defer rows.Close()

	//컬럼을 받아온다.
	cols, err := rows.Columns()

	//err발생했는지 확인한다.
	if err != nil {
		return nil, err
	}

	data := make([]interface{}, len(cols))

	for i, _ := range data {
		var d []byte
		data[i] = &d
	}

	results := make([]map[string]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(data...)
		if err != nil {
			return nil, err
		}
		result := make(map[string]interface{})
		for i, item := range data {
			result[cols[i]] = string(*(item.(*[]byte)))
		}
		results = append(results, result)
	}

	return results, nil
}

//SQL을 실행합니다.
func (dbms DBMS) MySQLExec(query string, db *sql.DB) (int64, int64, error) {
	defer db.Close()

	//SQL을 실행합니다.
	result, err := db.Exec(query)
	if err != nil {
		return -1, -1, err
	}

	// 변경된 row의 갯수를 가져옵니다.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}

	// 변경된 id를 가져옵니다.
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}

	return rowsAffected, lastInsertId, nil
}

// // CRUD
// func CreateTable(query string, table string, db *sql.DB, dbms *DBMS) {
// 	dbms.MySQLExecQuery(query, db)
// }

func Find(fields []string, table string, db *sql.DB, dbms *DBMS) []map[string]interface{} {
	results, _ := dbms.MySQLExecQuery(QuerySelect(fields, table), db)
	return results
}

func (dbms DBMS) InsertOne() {

}

func Insert(data []map[string]interface{}, table string, db *sql.DB, dbms *DBMS) {
	dbms.MySQLExec(QueryInsert(data, table), db)
}

// func Update(data map[string]interface{}, table string, db *sql.DB, dbms *DBMS) {
// 	dbms.MySQLExec(QueryUpdate(data, table), db)
// }

// func UpsertOne(data map[string]interface{}, table string, db *sql.DB, dbms *DBMS) {
// 	dbms.MySQLExec(QueryUpsertOne(data, table), db)
// }

// func Upsert(data []map[string]interface{}, table string, db *sql.DB, dbms *DBMS) {
// 	dbms.MySQLExec(QueryUpsert(data, table), db)
// }
