package main

import "fmt"

type Person struct {
	// properties
	firstName string
	lastName  string
	age       uint 
}

// انواع روش ساخت یک استراکت

func main() {
	//1
	person_1 := Person{firstName: "John", lastName: "lotfi", age: 19}

	//2 
	// ایجاد پرسن با این روش 
	// این یک اشاره گر است

	person_2 := new(Person)
	person_2.firstName = "John"
	person_2.lastName = "lotfi"
	person_2.age = 19

	//3
	// از جنس اشاره گر است
	person_3 := &Person{firstName: "John", lastName: "lotfi", age:19}

	//4
	// روش چهارم یکسری مزایا دارد می توانیم یکسری شرط بگذاریم مثلا اگر سن منفی وارد شد خطا بده	
	person_4 := NewPerson("John","lotfi",0)

	person_5 := NewPerson("jo" , "samadi" , 12)



	fmt.Println(person_1)
	fmt.Println(person_2)
	fmt.Println(person_3)
	fmt.Println(person_4)
	fmt.Println(person_5



}

func NewPerson(firstName string, lastName string , age uint) *Person { // uint ==> فقط مثبت می گیرد
	if age <= 0	{
		return nil
	} else if len(firstName) <3{
		return nil
	}
	return &Person{firstName: firstName, lastName: lastName , age: age}
}