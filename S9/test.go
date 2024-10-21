package main

import "fmt"

type Employee struct {
	FullName      string
	Age         int
	SalaryOfMonth int
}

func main() {

	employee1 := Employee{"Ali Ahmadi", 32, 2000}
	employee2 := new(Employee)
	employee2.FullName = "Vahid lotfi"
	employee2.Age  =18
	employee2.SalaryOfMonth = 3000



	fmt.Println(employee1)
	fmt.Println(CalcYearlyIncom(employee1))
	fmt.Println(employee1.calcYearIncomMethod())
	fmt.Println(employee1.CalcAgeToDay())


}


func CalcYearlyIncom(employee Employee) int{
	return employee.SalaryOfMonth*12
}

func (employee Employee) calcYearIncomMethod() int {
	return employee.SalaryOfMonth*12 
}

func (employee Employee) CalcAgeToDay() int {
	return employee.Age*365 
}