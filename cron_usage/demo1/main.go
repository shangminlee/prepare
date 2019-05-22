package main

import (
    "fmt"
    "github.com/gorhill/cronexpr"
    "time"
)

func main()  {
    var (
        expr     *cronexpr.Expression
        err      error
        now      time.Time
        nextTime time.Time
    )
    // crontab 表达式
    // 哪一分钟(0 - 59) 哪小时(0 - 23), 哪天(1 - 31), 那月(1 - 12), 星期几(0 - 6)

    // 每 1 分钟执行 1 次
    //if expr, err = cronexpr.Parse("* * * * *"); err != nil {
    //    panic(err)
    //    return
    //}

    // 每 5 分钟执行一次
    //if expr, err = cronexpr.Parse("*/5 * * * *"); err != nil {
    //    panic(err)
    //    return
    //}

    // 每 5 秒执行一次
    if expr, err = cronexpr.Parse("*/5 * * * * * * *"); err != nil {
        panic(err)
        return
    }

    // 当前时间
    now = time.Now()

    nextTime = expr.Next(now)

    time.AfterFunc(nextTime.Sub(now), func() {
        fmt.Println("执行调度任务 -->")
    })

    fmt.Printf(" current Time : %v \n next time : %v \n ", now, nextTime)

    time.Sleep(5 * time.Second)
}
