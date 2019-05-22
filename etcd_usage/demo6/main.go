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
        //delRes  *clientv3.DeleteResponse
        lease   clientv3.Lease
        putRes   *clientv3.PutResponse
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

    // 创建一个租约
    lease = clientv3.NewLease(client)
    leaseGrandResp, err := lease.Grant(context.TODO(), 5)
    leaseId := leaseGrandResp.ID

    contextTime, cancelFunc:= context.WithTimeout(context.TODO(), 6 * time.Second)
    keepAliveLease, err := lease.KeepAlive(contextTime, leaseId)

    // 消费自动续约
    go func() {
        timeOut := time.After(6 * time.Second)
        for {
            select {
            case keepAliveRes := <-keepAliveLease:
                if keepAliveRes == nil {
                    fmt.Println("停止续租")
                    time.Sleep(time.Second)
                    goto END
                }
                fmt.Println("续租ID", keepAliveRes.ID)

            case <- timeOut:
                fmt.Println("执行退出------>")
                cancelFunc()
                //return
                goto END
            }
        }
        END:
    }()

    key := "/cron/lock/job1"
    // 使用租约
    kv = clientv3.NewKV(client)
    putRes, err = kv.Put(context.TODO(),key,
        "",clientv3.WithLease(leaseId))

    fmt.Println(putRes)

    // 验证租约是否成功
    for {
        getRes, err = kv.Get(context.TODO(),key,)

        if len(getRes.Kvs) > 0 {
            fmt.Println(getRes.Kvs)
            time.Sleep(time.Second)
        }else {
            fmt.Println("租约失效了！")
            break
        }
    }

}
