package main

import "fmt"

type Person1 struct {
	Name   string
	family string
	Age    int
}

func main() {

	persons := make(map[int]Person1)
	persons[1] = Person1{Name: "vahid", family: "lotfi", Age: 32}
	persons[2] = Person1{Name: "vahid", family: "lotfi", Age: 32}

	fmt.Printf("persons list :%v ",persons)

}