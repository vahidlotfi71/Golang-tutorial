package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select{
		case <- ctx.Done():
			fmt.Printf("work stoped\n")
			return
		default :
			fmt.Print("work in progressing\n")
			time.Sleep(time.Microsecond *100)
		}
	}
}

func main() {
	ctx , cancel := context.WithCancel(context.Background())

	go doWork(ctx)

	time.Sleep(time.Microsecond *1000)
	cancel()

	time.Sleep(time.Second * 2)
	fmt.Println("done...")


}