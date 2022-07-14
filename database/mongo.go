// [Golang + MongoDB으로 CRUD 구현하기](https://soyoung-new-challenge.tistory.com/107)

// [https://flyingsquirrel.medium.com/go로-mongodb에-쿼리-날리기-da10a91aba33](https://flyingsquirrel.medium.com/)

// [MongoDB Go Driver Tutorial Part 1: Connecting, Using BSON, and CRUD Operations](https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial-part-1-connecting-using-bson-and-crud-operations)

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "fmt"
	// "time"
	// "context"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// // // get MongoDB Authorization info7
// func getAuth() m.Auth {
//     data, err := os.Open("[DB 접속정보.json 위치]")
//     U.CheckErr(err)

//     var auth m.Auth
//     byteValue, _ := ioutil.ReadAll(data)
//     json.Unmarshal(byteValue, &auth)

//     return auth
// }

// config := env.GetConfigYaml(YAML_NAME, serverName)  // map[string]interface{}

// ConnectDB to MongoDB
// func ConnectMongoDB(serverName string) (client *mongo.Client, ctx context.Context) {
func ConnectMongoDB(serverName string) (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	// Timeout 설정을 위한 Context생성
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	// ctx = context.Background()

	// 위에서 작성한 함수 사용
	// YAML_NAME := "database_conn"
	// serverName := "mongo_Oracle3_root_ex"
	// config := env.GetConfigYaml(YAML_NAME, serverName)

	// client, _ := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%v", config["host"], config["port"])))
	// Auth에러 처리를 위한 client option 구성
	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%v", config["host"], config["port"])).SetAuth(options.Credential{
	// 	Username: config["user"].(string),
	// 	Password: config["password"].(string),
	// })

	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%v", config["host"], config["port"]))
	// clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:3327")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// MongoDB 연결
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary()) // Primary DB에 대한 연결 체크
	if err != nil {
		log.Fatal(err)
	}

	// return client, ctx
	return client, ctx, cancel
}

func GetCollection(client *mongo.Client, colName, dbName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

func MgInsertOne(client *mongo.Client, colName, dbName string) {
	// collection := client.Database(dbName).Collection(colName)
	// insertResult, _ := collection.InsertOne(context.TODO(), bson.D{
	// 	{"userID", "test1234"},
	// 	{"array", bson.A{"flying", "squirrel", "dev"}},
	// })
	// insertResult, _ := client.Database(dbName).Collection(colName).InsertOne(context.TODO(), bson.D{
	// 	{Key: "title", Value: "The Polyglot Developer Podcast"},
	// 	{Key: "author", Value: "Nic Raboy"},
	// })
	insertResult, _ := client.Database(dbName).Collection(colName).InsertOne(context.TODO(), bson.D{
		{"a", 100},
		{"b", 200},
	})
	fmt.Println(insertResult)
}

func MgUpdateOne(client *mongo.Client, colName, dbName string) {
	filter := bson.D{{"a", 3}}

	update := bson.D{
		{"$inc", bson.D{
			{"b", 100},
		}},
	}

	updateResult, err := client.Database(dbName).Collection(colName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateResult)
}

// // CreateUser func
// func CreateUser(googleID, name string) string {
//     // DB 접속
//     client, ctx, cancel := ConnectDB()

//     // 함수 종료 뒤 연결을 끊어지도록 설정
//     defer client.Disconnect(ctx)
//     defer cancel()

//     // 필터 값 정의
//     filter := bson.M{"googleId": googleID, "name": name}

//     // DB에 값이 존재하는지 확인
//     num, err := GetCollection(client, "[collection이름]").CountDocuments(ctx, filter)
//     U.CheckErr(err)

//     // 새로 넣을 데이터 정의
//     newData := m.UserInfo{
//         GoogleID: googleID,
//         Name:     name,
//         Email:    email,
//     }

//     // DB값이 존재하지 않으면
//     if num == 0 {
//         _, err := GetCollection(client, "[collection이름]").InsertOne(ctx, newData)
//         U.CheckErr(err)
//     }

//     return "create!"
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	// 몽고DB 연결
// 	clientOptions := options.Client().ApplyURI("몽고 DB 주소를 입력해주세요")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("몽고 DB에 연결했습니다!")

// 	// 내용을 적을 부분

// 	// 몽고DB 연결 끊기
// 	uesrsCollection := client.Database("test").Collection("users")
// 	fmt.Println(uesrsCollection)

// 	err = client.Disconnect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("몽고DB 연결을 종료했습니다!")
// }
