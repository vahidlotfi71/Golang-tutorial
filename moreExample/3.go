package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "message sended from ch1🤣 "
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "message sended from ch2😂 "
	}()

	for i:=0 ; i<2 ; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
