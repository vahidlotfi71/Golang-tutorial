package main

import (
	"fmt"
	"time"
)

func printNumber(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("number: %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumber(3)
	fmt.Printf("این پیام هم زمان نمایش داده می شود\n")
	time.Sleep(2 * time.Second)

}