package main

import "fmt"

type Person struct {
	Name   string
	fmaily string
	age    int
}

func main() {

	person :=make(map[string]Person)

	person["121223231"] = Person{Name: "alice", age: 20, fmaily: "lotfield"}
	person["23456"] = Person{Name: "vahid", age: 20, fmaily: "lotfi"}


	fmt.Println(person)

	va := person["23456"]

	fmt.Print(va)

}
