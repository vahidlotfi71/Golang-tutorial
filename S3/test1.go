package main

import "fmt"

func main() {

	i, j := 12, 14
	var IP *int = &i
	var JP *int = &j

	fmt.Println("IP: ", *IP, " JP: ", *JP)
}
