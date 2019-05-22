package main

import (
    "context"
    "fmt"
    "os/exec"
    "strings"
    "time"
)

type result struct {
    output []byte
    err    error
}
func main()  {
    var (
        ctx        context.Context
        cancelFunc context.CancelFunc
        cmd        *exec.Cmd
        resultChan chan *result
        res        *result
    )
    resultChan = make(chan *result)
    ctx, cancelFunc = context.WithCancel(context.TODO())
    go func (){
        var (
            output []byte
            err    error
        )
        cmd = exec.CommandContext(ctx ,"/bin/bash","-c", "sleep 4; echo hello;")

        if output, err = cmd.CombinedOutput(); err != nil {
            //panic(err)
        }
        resultChan <- &result{output:output,err:err}
    }()

    time.Sleep(1 * time.Second)

    cancelFunc()
    res = <- resultChan
    fmt.Println(strings.Trim(string(res.output),"\n"), res.err)
}
