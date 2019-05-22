package main

import (
    "context"
    "fmt"
    "go.etcd.io/etcd/clientv3"
    "time"
)

func main() {
    var (
        config  clientv3.Config
        client  *clientv3.Client
        err     error
        kv      clientv3.KV
    )

    // 客户端配置
    config = clientv3.Config{
        Endpoints:[]string{"127.0.0.1:2379"},
        DialTimeout: 5 * time.Second,
    }

    // 监理连接
    if client, err = clientv3.New(config); err != nil {
        panic(err)
        return
    }

    kv = clientv3.NewKV(client)

    key := "/cron/jobs/job8"
    // 创建 Op : operation

    // Op put
    putOp := clientv3.OpPut(key, "hello operation")

    opRes, err := kv.Do(context.TODO(), putOp)

    fmt.Println("写入 Revision : ", opRes.Put().Header.Revision)

    // Op get
    getOp := clientv3.OpGet(key)

    opRes, err = kv.Do(context.TODO(), getOp)

    fmt.Println("读取 Revision : ", opRes.Get().Kvs[0].ModRevision)
}
