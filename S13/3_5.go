package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	GetEmployeeInfo("ali", "samadi",4567899)

	fmt.Println("End of main function")

}

func GetEmployeeInfo(firstName, lastName string, salary int) float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error : ", err)
			fmt.Println("Recovering error...")
			debug.PrintStack()
		}
	}()
	if firstName == "" {
		panic("First name is empty")
	}
	fmt.Println("firstName", firstName)

	if lastName == "" {
		panic("Last name is empty")
	}
	fmt.Println("lastName", lastName)

	if salary <= 0 {
		panic("salary is less than 0")
	}
	fmt.Println("salary" , salary)

	return calculateTax(salary)
}

func calculateTax(salary int) float64 {
	tax := float64(salary) * 0.10
	if tax > 1000 {
		panic("Tax is greater than 1000")
	}
	fmt.Println("tax: " , tax)
	return tax
}