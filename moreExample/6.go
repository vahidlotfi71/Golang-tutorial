package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func calculate(number []int , wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range number {
		numpow := math.Pow(float64(num), 3)
		fmt.Printf("number :%d , numpow: %.2f\n", num, numpow)
	}
}

func main() {
	runtime.GOMAXPROCS(4	)
	wg := sync.WaitGroup{}
	numberList1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	numberList2 := []int{10, 11, 12, 13, 14, 15, 16, 17}

	wg.Add(2)
	go calculate(numberList1 , &wg)
	go calculate(numberList2 , &wg)
	wg.Wait()

	fmt.Print("\nfinished...")
}
