package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	Print("Hello")
	Print(123)
	Print(true)
	Print([]int{1, 2, 3})
	Print(map[string]int{"a": 1, "b": 2})
	Print(Person{"Ali"})

}

func Print(input interface{}) {
	switch input.(type) {
	case string :
		fmt.Println("String:", input)
	case int:
		fmt.Println("Int:", input)
	case bool:
		fmt.Println("Bool:", input)
	case []int:
		fmt.Println("Array:", input)
	case map[string]int:
		fmt.Println("Map:", input)
	case Person:
		fmt.Println("Person:" , input)

}
}
