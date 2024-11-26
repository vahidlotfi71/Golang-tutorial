package main

import "fmt"

type Person2 struct {
	Name string `json:"name"`
}

func main() {

	chap(1)
	chap(12.12)
	chap([]int{1, 2, 3, 4, 5, 6})
	chap(map[int]string{1: "one", 2: "two", 3: "three"})
	chap(true)

}

func chap(input interface{}) {
	switch input.(type) {
	case string:
		fmt.Println("string:", input)
	case int:
		fmt.Println("int:", input)
	case float64:
		fmt.Println("float64: ", input)
	case bool:
		fmt.Println("bool:", input)
	case Person2:
		fmt.Println("Person2: ", input)
	case []int:
		fmt.Println("slice: ", input)
	case map[int]string:
		fmt.Println("map: ", input)
	}

}
