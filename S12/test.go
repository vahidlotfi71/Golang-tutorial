package main

import "fmt"

type Number22 interface {
	int | float64
}

func main() {
	mySlic1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	mySlic2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	println(SumSlic(mySlic1))
	fmt.Println(SumSlic(mySlic2))

}

func SumSlic[T Number22](slices []T) (sum T) {
	for _, v := range slices {
		sum += v
	}
	return
}
