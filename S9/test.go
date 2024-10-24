package main

import "fmt"

type Person1 struct {
	int
	string
	float64
}

func main() {

	person1 := struct {
		NationaCode int
		Name        string
		Age         int
	}{
		NationaCode: 2960220481,
		Name:        "vahid",
		Age:         32,
	}

	person2 := Person1{121212121, "vahid", 20.11}

	fmt.Println(person1)
	fmt.Println(person2)
}