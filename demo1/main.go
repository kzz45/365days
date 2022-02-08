package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	var cmd *exec.Cmd
	var res *result
	var ctx context.Context
	var cancelFunc context.CancelFunc
	var resultChan chan *result

	resultChan = make(chan *result, 100)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var err error
		var output []byte
		cmd = exec.CommandContext(ctx, "bash", "-c", "sleep 2; echo fuck!")
		output, err = cmd.CombinedOutput()
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)
	cancelFunc()

	res = <-resultChan

	fmt.Println(res.err, string(res.output))
}
