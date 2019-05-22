package main

import (
    "go.etcd.io/etcd/clientv3"
    "time"
)

func main() {
    var (
        config  clientv3.Config
        client  *clientv3.Client
        err     error
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


    client = client
}
