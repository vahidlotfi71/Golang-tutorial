package main

import (
	"fmt"
	// "runtime"
	"sync"
	// "time"
)

func PrintNumber(name string , wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Printf("%s , %d \n", name, i)

	}
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}

	wg.Add(2)
	go PrintNumber("Goroutine one" , &wg)
	go PrintNumber("Gorutine towe", &wg)
	wg.Wait()
}
