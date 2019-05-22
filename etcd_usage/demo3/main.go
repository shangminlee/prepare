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
        getRes  *clientv3.GetResponse
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
    kv = clientv3.NewKV(client)

    if putRes, err := kv.Put(context.TODO(),"cron/jobs/job1","hello, world!",clientv3.WithPrevKV()); err != nil {
        panic(err)
    }else {
        fmt.Printf("putRes : %v \n", putRes)
        if putRes.PrevKv != nil {
            fmt.Println("PreValue", string(putRes.PrevKv.Value))
        }
    }

    if putRes, err := kv.Put(context.TODO(),"cron/jobs/job2","hello, world!",clientv3.WithPrevKV()); err != nil {
        panic(err)
    }else {
        fmt.Printf("putRes : %v \n", putRes)
        if putRes.PrevKv != nil {
            fmt.Println("PreValue", string(putRes.PrevKv.Value))
        }
    }

    if putRes, err := kv.Put(context.TODO(),"cron/jobs/job3","hello, world!",clientv3.WithPrevKV()); err != nil {
        panic(err)
    }else {
        fmt.Printf("putRes : %v \n", putRes)
        if putRes.PrevKv != nil {
            fmt.Println("PreValue", string(putRes.PrevKv.Value))
        }
    }

    getRes, err = kv.Get(context.TODO(),"cron/jobs/job1")
    if err != nil {
        panic(err)
    }
    for i, kypair := range getRes.Kvs {
        fmt.Println(i," : ", string(kypair.Key), string(kypair.Value))
    }

}
