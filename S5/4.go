package main

import "fmt"

func main() {

	var num int
	var provinceName string

	print("Please enter num: ")
	fmt.Scanln(&num)

	switch num {
	case 10,20,30,40,50,60,70,80,90,11,22,33,44,55,66,77,88,99:
		provinceName = "Tehran"
	case 17:
		provinceName = "A.garbi"
	case 15:
		provinceName = "Tabriz"
	default:
		provinceName = "Unknown"
}

	fmt.Println(provinceName) 
}

 