package example

import (
	"encoding/json"
)

type Persons struct {
	Name   string `json:"name"`
	Family string `json:"-"`  // اگر دش بگزاریم دیگر در خروجی سن را نمایش 
	Age    int	  `json:"age ,omitempty"` // omitempty اگر مقدارش صفر بود دیگر در خروجی نیاور
}

func Marshal1() {
	person1 := Persons{Name: "John", Family: "lotfi", Age: 32}
	person2 := Persons{Name: "Ali", Family: "samadi", Age: 45}
	person3 := Persons{Name: "mohammad", Family: "ahmadi", Age:0}

	person1Json, err := json.Marshal(person1)
	if err != nil {
		panic(err)
	}

	person2Json , err := json.Marshal(person2)
	if err != nil {
		panic(err)
	}

	person3Json , err := json.Marshal(person3)
	if err != nil {
		panic(err)

	}

	println(string(person1Json))
	println(string(person2Json))
	println(string(person3Json))
}


func Marshal2() {
	person1 := Persons{Name: "John", Family: "lotfi", Age: 32}
	person2 := Persons{Name: "Ali", Family: "samadi", Age: 45}
	person3 := Persons{Name: "mohammad", Family: "ahmadi", Age:0}

	personsList := []Persons{person1, person2, person3}
	person1Json, err := json.Marshal(personsList)
	if err != nil {
		panic(err)
	}
	println(string(person1Json))
}