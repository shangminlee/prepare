package main

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

// 任务执行 开始 结束 时间
type TimePoint struct {
    StartTime int64
    EndTime   int64
}

// 日志记录
// Golang   打标签 `json: "jobName"`
// MoongoDB 打标签 `bson: "jobName"`
type LogRecord struct {
    JobName   string `bson:"jobName"`
    Command   string `bson:"command"`
    Err       string `bson:"err"`
    Content   string `bson:"content"`
    TimePoint TimePoint `bson:"timePoint"`
}

type FindByJobName struct {
    JobName string `bson:"jobName"`
}

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
    database := client.Database("cron")


    // 3. 选择表
    collection := database.Collection("log")

    cond := &FindByJobName{JobName:"job10"}

    // 查询条件
    var findOps = options.FindOptions{}
    cursor, err := collection.Find(context.TODO(), cond, findOps.SetSkip(0), findOps.SetLimit(10))
    if err != nil {
        return
    }
    defer func() {
        err := cursor.Close(context.TODO())
        fmt.Println(err)
    }()

    for cursor.Next(context.TODO()) {
        var result LogRecord
        err := cursor.Decode(&result)

        if err != nil {
            panic(err)
        }

        fmt.Println(result)
    }

}
