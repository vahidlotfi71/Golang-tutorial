package main

import "fmt"

func sendData(ch chan string) {
	ch <- "send data from Goroutine"
}

func main() {
	ch1 := make(chan string)
	go sendData(ch1)
	fmt.Printf(<-ch1)
}