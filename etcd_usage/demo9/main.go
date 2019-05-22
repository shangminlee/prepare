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

    // 乐观锁
    // 1. 上锁 (创建租约， 自动续约， 拿着租约去抢占一个 Key)
    lease := clientv3.NewLease(client)

    leaseGrant, err := lease.Grant(context.TODO(), 10)
    if err != nil {
        panic(err)
    }
    leaseId := leaseGrant.ID

    // 创建可取消的Context, 用于取消
    ctx, cancelFunc := context.WithCancel(context.TODO())
    defer cancelFunc() // 最终取消 续约
    defer func() {
        _, err := lease.Revoke(context.TODO(), leaseId)
        if err != nil {
            fmt.Println("Revoke lease Error")
        }
    }()

    keepAliveChan, err := lease.KeepAlive(ctx,leaseId)
    if err != nil {
        panic(err)
    }
    go func() {
        for {
            select {
            case keepAliveRes := <- keepAliveChan:
                if keepAliveRes != nil {
                    fmt.Println("收到租约应答", keepAliveRes.ID)
                }else{
                    fmt.Println("租约已经失效")
                    goto END
                }
            }
        }
        END:
    }()

    // 抢占Key
    key := "/cron/lock/job9"

    // 创建KV
    kv  = clientv3.NewKV(client)

    // 创建事务
    txn := kv.Txn(context.TODO())

    // 如果Key 不存在
    txn.If(
        clientv3.Compare(clientv3.CreateRevision(key), "=", 0),
    ).Then(
        clientv3.OpPut(key, "lock value", clientv3.WithLease(leaseId)),
    ).Else(
        clientv3.OpGet(key),
    )

    txnRes, err := txn.Commit()
    if err != nil {
        panic(err)
    }

    // 如果没有抢占成功
    if !txnRes.Succeeded {
        fmt.Println("锁被占用 ： ", string(txnRes.Responses[0].GetResponseRange().Kvs[0].Value))
        return
    }

    // 2. 处理业务
    fmt.Println("开始处理业务")
    time.Sleep(60 * time.Second)

}
