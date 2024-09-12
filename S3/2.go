package main

import "fmt"






func maine() {

	var num int
	var num8 int8
	var num16 int16
	var num32 int32
	var num64 int64

	fmt.Println("num %d bytes \n" , unsafe.Sizeof(num))
	fmt.Println("num %d bytes \n" , unsafe.Sizeof(num8))
	fmt.Println("num %d bytes \n" , unsafe.Sizeof(num16))
	fmt.Println("num %d bytes \n" , unsafe.Sizeof(num32))
	fmt.Println("num %d bytes \n" , unsafe.Sizeof(num64))

	var a = bite.unitsize
	fmt.Println(a)


}

