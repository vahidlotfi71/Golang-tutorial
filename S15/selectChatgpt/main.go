// package main

// import (
// 		"fmt"
// )

// func main() {

// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {

// 		ch1 <- "send data from ch1"
// 	}()

// 	go func() {

// 		ch2 <- "send data from ch2"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println("data received from ",msg1)

// 		case msg2 := <-ch2:
// 			fmt.Println("data received from ",msg2)
// 		default:
// 			fmt.Println("there isn't any ready channel")

// 		}
// 	}
// }