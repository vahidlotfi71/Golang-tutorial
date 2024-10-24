package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string  `json:"first_name"`
	Age       int	   
	Family    string  `json:"last_name"`
	Gender    string
	Height    int
	Weight    int
	HairColor string
}

func main() {
	person1 := Person{Name: "John", Age: 32, Family: "lotfi", Gender: "male", Height: 173, Weight: 70, HairColor: "Black"}

	fmt.Println(person1)

	personJson , _ := json.MarshalIndent(person1, "", " ")

	fmt.Println(string(personJson))
}