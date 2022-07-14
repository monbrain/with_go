package database

import (
	// "database/sql"
	"fmt"
	"strings"
	// "gopkg.in/yaml.v3"
	// // "time"
	// _ "github.com/go-sql-driver/mysql"
	// "with_go/env"
)

func wrapWithQuotationMark[T any](val T) string {
	// val.(type)
	ret := fmt.Sprintf("%v", val)
	if fmt.Sprintf("%T", val) == "string" {
		ret = `"` + ret + `"`
	}
	return ret
}

// func KeysOfMap(m interface{}) (keys []interface{}) {
// 	v := reflect.ValueOf(m)
// 	if v.Kind() != reflect.Map {
// 		fmt.Errorf("input type not a map: %v", v)
// 	}

// 	for _, k := range v.MapKeys() {
// 		keys = append(keys, k.Interface())
// 	}
// 	return keys
// }

func KeysValuesOfMap(data map[string]interface{}) (keys []string, vals []string) {
	for key, val := range data {
		keys = append(keys, key)
		vals = append(vals, wrapWithQuotationMark(val))
	}
	return keys, vals
}

func ValuesStringByKeys(data map[string]interface{}, keys []string) string {
	vals := []string{}
	for _, key := range keys {
		vals = append(vals, wrapWithQuotationMark(data[key]))
	}
	return "(" + strings.Join(vals, ", ") + ")"
}

func QuerySelect(fields []string, table string) string {
	keys := "*"
	if len(fields) > 0 {
		keys = strings.Join(fields, ", ")
	}
	return "SELECT " + keys + " FROM " + table + ";"
}

func QueryInsertOne(data map[string]interface{}, table string) string {
	keys, vals := KeysValuesOfMap(data)
	return "INSERT INTO " + table + " (" + strings.Join(keys, ", ") + ")" + " VALUES (" + strings.Join(vals, ", ") + ");"
}

func QueryInsert(data []map[string]interface{}, table string) string {
	keys, _ := KeysValuesOfMap(data[0])
	vals := []string{}
	for _, d := range data {
		vals = append(vals, ValuesStringByKeys(d, keys))
	}
	return "INSERT INTO " + table + " (" + strings.Join(keys, ", ") + ")" + " VALUES " + strings.Join(vals, ", ") + ";"
}

func QueryUpdate(data map[string]interface{}, table string) string {
	keys, vals := KeysValuesOfMap(data)
	query := "UPDATE " + table + " SET "
	for i := range keys {
		query += keys[i] + " = " + vals[i] + ", "
	}
	return query[:len(query)-2]
}

func QueryUpsert(data map[string]interface{}, table string) string {
	// keys, vals := KeysValuesOfMap(data)
	// query := "UPDATE " + table + " SET "
	// for i := range keys {
	// 	query += keys[i] + " = " + vals[i] + ", "
	// }
	return ""
}

// // func StrInsertKeyValues(data map[string]interface{}) string {
// // 	keys, vals := KeysValuesOfMap(data)
// // 	return "(" + strings.Join(keys, ", ") + ")"  strings.Join(vals, ", ")
// // }

// // func setInsertKeyValsFromMaps(data []map[string]interface{}) (string, string) {
// // 	keys := keySlice(data[0])
// // 	vals := []string{}
// // 	for _, d := range data {
// // 		vals = append(vals, valSql(d, keys))
// // 	}

// // 	return "(" + strings.Join(keys, ", ") + ")", strings.Join(vals, ", ")
// // }

// func main() {
// 	data := []map[string]interface{}{
// 		{"a": 1, "b": "cd"},
// 		{"a": 3, "b": "ef"},
// 	}

// 	// Insert
// 	fmt.Println(QueryInsert(data, "table1"))
// 	// Update
// 	fmt.Println(QueryUpdate(data[0], "table1"))

// }
