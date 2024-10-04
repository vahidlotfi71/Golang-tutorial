package main

import "fmt"

type Person struct {
	name   string
	family string
	age    int
}

func main() {
	persons := make(map[string]Person)

	// با ایجاد اسلایس و افزودن کد ملی ها ترتیب نمایش را درست می کنیم
	personslices := []string{}
	persons["2960220481"] = Person{name: "vahid", family: "lotfy", age: 32}
	personslices = append(personslices,"2960220481")
	persons["2960229988"] = Person{name: "hasan", family: "mehtari", age: 28}
	personslices = append(personslices,"2960229988")
	persons["1566658985"] = Person{name: "maman", family: "mosvi", age: 57}
	personslices = append(personslices,"1566658985")
	persons["9896265984"] = Person{name: "sara", family: "hasani", age: 23}
	personslices = append(personslices,"9896265984")
	persons["3216549858"] = Person{name: "sahar", family: "rahmani", age: 5}
	personslices = append(personslices,"3216549858")
	persons["2165498512"] = Person{name: "reza", family: "mosvi", age: 7}
	personslices = append(personslices,"2165498512")
	persons["6519879819"] = Person{name: "talan", family: "nosrati", age:19}
	personslices = append(personslices,"6519879819")
	persons["9846512165"] = Person{name: "rahmat", family: "mosvi", age: 50}
	personslices = append(personslices,"9846512165")
	persons["6541621989"] = Person{name: "ramazan", family: "zamani", age: 42}
	personslices = append(personslices,"6541621989")


	fmt.Printf("len: %d\n", len(persons))


	// ترتیب نمایش دیتا در خروجی بهم می خورد
	for _  , person := range persons{
		fmt.Printf("%v\n", person)
	}

	for _ , person := range personslices{
		fmt.Printf("%v\n", persons[person])
	}

}