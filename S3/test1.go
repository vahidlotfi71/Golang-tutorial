package main

import "fmt"

func main() {

	a := "vahid"
	b := "hasan"
	c := 123

	mystr := "hi vahid wellcome to learn GO language"

	fmt.Printf("a: %v ,Type: %T \n", a, a)
	fmt.Printf("b: %v ,Type: %T\n", b, b)
	fmt.Printf("c: %v ,Type: %T\n", c, c)
	println()
	fmt.Printf("mystr: %v , Type: %T , len: %d \n", mystr, mystr,len(mystr))
}