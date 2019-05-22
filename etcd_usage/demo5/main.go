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
        delRes  *clientv3.DeleteResponse
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

    getRes, err =kv.Get(context.TODO(),"cron/jobs", clientv3.WithFromKey())

    fmt.Println(getRes.Kvs)

    delRes, err = kv.Delete(context.TODO(), "cron/jobs", clientv3.WithFromKey())

    fmt.Println(delRes.Deleted)

    getRes, err =kv.Get(context.TODO(),"cron/jobs", clientv3.WithFromKey())

    fmt.Println(getRes.Kvs)

}
