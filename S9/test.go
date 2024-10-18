package main

import "fmt"

type Person struct {
	name   string
	family string
	age    uint
}

func main() {

	person_1 := Person{"vahid", "lotfi", 32}
	person_2 := &Person{name: "vahid", family: "lotfi", age: 32}
	person_3 := NewPerson("vahid", "lotfi",32)

	person_4 := new(Person)
	person_4.name = "vahid"
	person_4.family = "vahid"
	person_4.age  = 32


	fmt.Println(person_1)
	fmt.Println(person_2)
	fmt.Println(person_3)
	fmt.Println(person_4)
}


func NewPerson(name string, family string, age uint) *Person {

	return &Person{name: name, family: family, age: age}}
