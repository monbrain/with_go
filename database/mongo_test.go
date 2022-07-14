package database_test

import (
	"fmt"
	// "log"
	"testing"
	"with_go/database"
	// "with_go/env"
)

// func TestConnectDb(t *testing.T) {
// 	client, ctx, cancel := database.ConnectMongoDB("mongo_Oracle3_root_ex")
// 	// client, ctx, cancel := database.ConnectMongoDB("mongo_HMS_root_in")
// 	if client == nil {
// 		t.Error("Wrong result")
// 	}
// 	fmt.Println(client, ctx, cancel)
// }

func TestGetCollection(t *testing.T) {
	// client, ctx, cancel := database.ConnectMongoDB("mongo_Oracle3_root_ex")
	client, ctx, cancel := database.ConnectMongoDB("mongo_HMS_root_in")
	// client, ctx := database.ConnectMongoDB("mongo_HMS_root_in")
	if client == nil {
		t.Error("Wrong result")
	}
	// fmt.Println(ctx)
	fmt.Println(ctx, cancel)
	fmt.Println(database.GetCollection(client, "test_", "sats"))
}

// func TestInsertOne(t *testing.T) {
// 	client, ctx, cancel := database.ConnectMongoDB("mongo_HMS_root_in")
// 	if client == nil {
// 		t.Error("Wrong result")
// 	}
// 	fmt.Println(ctx, cancel)
// 	database.MgInsertOne(client, "test_", "sats")
// }

// func TestMgUpdateOne(t *testing.T) {
// 	client, ctx, cancel := database.ConnectMongoDB("mongo_HMS_root_in")
// 	if client == nil {
// 		t.Error("Wrong result")
// 	}
// 	fmt.Println(ctx, cancel)
// 	database.MgUpdateOne(client, "test_", "sats")
// }
