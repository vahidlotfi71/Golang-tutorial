package main

import (
	"context"
	"time"
)

func doWork(ctx context.Context){
	select {
	case <- time.After(time.Second *1):
		println("work done")
	case <- ctx.Done():
		println("time out ")
	}
}
func main() {
	ctx , cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	go doWork(ctx)

	time.Sleep(time.Second *4)
	println("process done")
}