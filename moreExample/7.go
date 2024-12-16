package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		ch1 <- 9385932919
		ch2 <- "number of vahid lotfi"
	}()

	msg1 := <-ch1
	msg2 := <-ch2




	fmt.Printf("msg1: %v \nmsg2: %v\n",msg1,msg2)

}
