package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan float64)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- 324256
			ch2 <- "vahid lofi"
			ch3 <- 343.3
		}
	}()

	select {
	case msg1 := <-ch1:
		fmt.Printf("msg1: %v\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("msg2: %v\n", msg2)
	case msg3 := <-ch3:
		fmt.Printf("msg3: %v\n", msg3)
	default:
		fmt.Printf("don't recive message from server")

	}

}
