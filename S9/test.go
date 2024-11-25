package main

import (
	"fmt"

)

type Animal struct {
	spatiouse string
	height    float64
	weight    float64
}

func main() {

	human := struct{
		Name string
		Age int
		height float64
		weight float64
	}{
		Name: "vahid",
		Age: 31,
		height: 173,
		weight: 70,

	}
	animal := Animal{spatiouse: "reptails" , height: 220 , weight: 323}

	fmt.Printf("%+v\n", human)
	fmt.Printf("%+v\n", animal)
}