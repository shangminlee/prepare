package main

import (
    "fmt"
    "github.com/gorhill/cronexpr"
    "time"
)

type CronJob struct {
    expr     *cronexpr.Expression
    nextTime time.Time
}

func main()  {
    var (
        cronJob       *CronJob
        expr          *cronexpr.Expression
        now           time.Time
        scheduleTable map[string]*CronJob // 任务调度表
    )

    scheduleTable = make(map[string]*CronJob)
    // 当前时间
    now = time.Now()

    expr, _ = cronexpr.Parse("*/5 * * * * * *")
    cronJob = &CronJob{
        expr:     expr,
        nextTime: expr.Next(now),
    }

    scheduleTable["job1"] = cronJob

    fmt.Printf("job1 --> %v \n",  cronJob)

    expr, _ = cronexpr.Parse("*/5 * * * * * *")
    cronJob = &CronJob{
        expr:     expr,
        nextTime: expr.Next(now),
    }

    scheduleTable["job2"] = cronJob
    fmt.Printf("job2 --> %v \n",  cronJob)

    go func() {
        for {
            now = time.Now()
            for jobName, cronJob := range scheduleTable {
                if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
                    go func(jobName string) {
                        fmt.Println("执行", jobName)
                    }(jobName)
                }
                cronJob.nextTime = cronJob.expr.Next(now)
            }
            select {
            case <- time.NewTimer(1 * time.Second).C:
                fmt.Println("----> next schedule ---->")
            }
        }
    }()

    time.Sleep(10 * time.Second)
}