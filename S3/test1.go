package main

import "fmt"

func main() {
	var unm int
	var num8 int8

	fmt.Println("num %d bytes", unsafe.Sizeof(num))
	fmt.Println("num8 %d bytes", unsafe.Sizeof(num8))

}