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
        //kv      clientv3.KV
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
    //client.Put()
    // 键值对
    //kv = clientv3.NewKV(client)

    if putRes, err := client.Put(context.TODO(),"cron/jobs/job1","hello, world!",clientv3.WithPrevKV()); err != nil {
        panic(err)
    }else {
        fmt.Printf("putRes : %v \n", putRes)
        if putRes.PrevKv != nil {
            fmt.Println("PreValue", string(putRes.PrevKv.Value))
        }
    }

}
