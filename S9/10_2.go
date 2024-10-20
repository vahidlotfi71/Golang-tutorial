package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	Family    string
	Gender    string
	Height    int
	Weight    int
	HairColor string
}

type PersonOptions func(*Person)  // تایپ این یک فانکشن است قرار پرسن را به عنوان ورودی بگیرد و یکسری مقادیر را برایش ست کن


func main() {

	person := NewPerson(SetName("Maryam"), SetFamily("Rezaei"), SetAge(23), SetGender("Female"), SetHeight(175), SetWeight(70), SetHairColor("Black"))

	fmt.Printf("Person : %+v" , person)

}

func NewPerson(options ...PersonOptions) *Person {  // می خواهیم پرسن اپشن هارا بگیرد و در نهایت یک پرسن برای من برگردانند
	person := &Person{}
	for _, option := range options {
		option(person)
	}
	return person
}

func SetName(name string) PersonOptions { // چون این از جنس فانکشن است نیازی نیست از استار استفاده کنیم
	return func(person *Person) { // یک پرسن از جنس پرسن میگیرد
		person.Name = name
	}
}

func SetAge(age int) PersonOptions {
	return func(person *Person) {
		person.Age = age
	}
}

func SetFamily(family string) PersonOptions {
	return func(person *Person) {
		person.Family = family
	}
}

func SetGender(gender string) PersonOptions {
	return func(person *Person) {
		person.Gender = gender
	}
}

func SetHeight(height int) PersonOptions {
	return func(person *Person) {
		person.Height = height
	}
}

func SetWeight(weight int) PersonOptions {
	return func(person *Person) {
		person.Weight = weight
	}
}

func SetHairColor(hairColor string) PersonOptions {
	return func(person *Person) {
		person.HairColor = hairColor
	}
}