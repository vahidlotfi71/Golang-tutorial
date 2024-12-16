package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

func printInt(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("number: %d\n", number)
}

func main() {

	var wg = sync.WaitGroup{}

	message := []string{"vahid", "hasan", "ali", "fatemeh", "reza"}
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, msg := range message {
		wg.Add(1)
		go printMessage(msg, &wg)
	}
	time.Sleep(time.Second * 1)
	for _, number := range numberList {
		wg.Add(1)
		time.Sleep(time.Microsecond *10)
		go printInt(number, &wg)
	}

	wg.Wait()
	fmt.Println("\n\nfinished😁👍")
}
