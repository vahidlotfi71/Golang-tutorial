package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 22
	ch <- 33

	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
