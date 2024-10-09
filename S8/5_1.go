package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{12, 32, 45, 32, 67, 78, 22, 1, 2, 68, 0, 32, 11}
	numbers1 := []int{12, 32, 45, 32, 67, 78, 22, 1, 2, 68, 0, 32, 11}

	fmt.Print("%v\n", numbers)
	fmt.Print("%v\n", numbers1)



	sort.Slice(numbers, func(i , j int ) bool { // یک اسلایس میگیرد و مرتب می کند 
		return numbers[i] > numbers[j] // براساس ><= خروجی را مرتب می کند 
	})


	sortingfunc := func(i , j int ) bool { // یک اسلایس میگیرد و مرتب می کند 
		return numbers1[i] < numbers1[j]
	}

	sort.Slice(numbers1, sortingfunc)	


	fmt.Print("%v\n", numbers)
	fmt.Print("%v\n", numbers1)

}