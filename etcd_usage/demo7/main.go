package main

import (
    "context"
    "fmt"
    "go.etcd.io/etcd/clientv3"
    "go.etcd.io/etcd/mvcc/mvccpb"
    "math/rand"
    "strconv"
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

    kv = clientv3.NewKV(client)

    resChan := make(chan int)
    key := "cron/jobs/job7"
    go func() {
        for {
            seed := strconv.Itoa(rand.Intn(1000))
            _, err := kv.Put(context.TODO(),key,"hello, world!" + seed,clientv3.WithPrevKV())

            //fmt.Println("Put ",putRes)

            _, err = kv.Delete(context.TODO(), "cron/jobs/job7")

            //fmt.Println("Del ",delRes)
            time.Sleep(time.Second)


            if err != nil {
                break
            }
        }
        resChan <- 1
    }()

    if getRes, err = kv.Get(context.TODO(),key); err != nil {
         fmt.Println("Error")
        return
    }
    if len(getRes.Kvs) > 0 {
        fmt.Println("当前值", string(getRes.Kvs[0].Value))
    }

    watchStartVersion := getRes.Header.Revision + 1;

    fmt.Println("从", watchStartVersion, "开始监听")

    watch := clientv3.NewWatcher(client)

    watchChan := watch.Watch(context.TODO(),
        key, clientv3.WithRev(watchStartVersion))


    timeOut := time.After(time.Minute)
    for {

        select {
        case wRes := <-watchChan:
            for _, event := range wRes.Events {
                switch event.Type {
                case mvccpb.DELETE:
                    fmt.Println("删除：",string(event.Kv.Key))
                case mvccpb.PUT:
                    fmt.Println("修改为：", string(event.Kv.Value), " Reversion", event.Kv.CreateRevision, event.Kv.Version, event.Kv.ModRevision)

                }
            }

        case <- timeOut:
            fmt.Println("退出监听")
            goto OUT

        }
    }
    OUT:
    <- resChan
    
}
