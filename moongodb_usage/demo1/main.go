package main

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "time"
)

func main() {

    // 1. 建立连接

    ctx, _ := context.WithTimeout(context.TODO(), 10 * time.Second)

    client, err := mongo.Connect(
        ctx,
        options.Client().ApplyURI("mongodb://localhost:27017"),
        options.Client().SetConnectTimeout(10 * time.Second),
        )

    if err != nil {
        panic(nil)
    }
    // 2. 选择数据库
    database := client.Database("my_db")


    // 3. 选择表
    collection := database.Collection("my_collection")

    fmt.Println(collection.Name())

}
