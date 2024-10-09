package main

import "fmt"

//به فانکشنن هایی که درون یه فانشن دیگر تعریف می شود این لاین فانکشن هم گفته می شود 

func main() {

	func() {
		fmt.Println("Hello world!")
	}() // با پارانتز تابع را کال می کنیم


	my_func1 := func() {
		fmt.Println("Hello world!")
	}
	my_func1()
	my_func1()

	println(func(numbers ... int) int {
		var total int
		for _ ,number := range numbers {
			total += number
	}
	return total
}(1,2,3,4,5))

	sum := func (numbers ... int) int {
		var total int 
		for _ , number := range numbers{
			total += number
		}
		return total
	}

	println(sum(1,2,3,4,5,6,7))
	println(sum(10,20,30,40,50,60,70))



}