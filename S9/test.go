package main

import "fmt"

const (
	SalaryBase            = 10000
	ExtraHourRateFullTime = 10
	ExtraHourRatePartTime = 10
	TaxRate               = 0.10
)

func main() {

	fullTimeEmployees := []FullTimeEmployee{
		{Id: 1, NationalCode: "1234567890", FullName: "Pejman Rezaee", ExtraHours: 40},
		{Id: 2, NationalCode: "4836524125", FullName: "Maryam Hosseini", ExtraHours: 120},
	}

	partTimeEmployees := []PartTimeEmployee{
		{Id: 3, NationalCode: "6563453455", FullName: "Milad Hassani", PartTimeHours: 100},
		{Id: 4, NationalCode: "5435435435", FullName: "Maryam Rezaee", PartTimeHours: 87},
	}

	var employye FullTimeEmployee = fullTimeEmployees[0]

	var employee1 PartTimeEmployee = partTimeEmployees[0]

	salary , tax := employye.SalaryCalculate(employye.ExtraHours)
	salary1 , tax1 := employee1.SalaryCalculate(employee1.PartTimeHours)

	fmt.Println("Employee : %v , salary : %f , tax : %f", employye , salary , tax)
	fmt.Println("Employee : %v , salary : %f , tax : %f", employee1 , salary1 , tax1)

	



}

type Employee interface{
	SalaryCalculate(hours float64)  (salary , tax float64)
}

type FullTimeEmployee struct {
	Id           int
	NationalCode string
	FullName     string
	Type         string
	ExtraHours   float64
}

type PartTimeEmployee struct {
	Id            int
	NationalCode  string
	FullName      string
	PartTimeHours float64
}

func (employee FullTimeEmployee) SalaryCalculate(extraHours float64) (salary, tax float64) {
	salary = SalaryBase + extraHours*ExtraHourRateFullTime
	tax = salary * TaxRate
	return
}

func (employee PartTimeEmployee) SalaryCalculate(extraHours float64) (salary, tax float64) {
	salary = ExtraHourRatePartTime * extraHours
	tax = salary * TaxRate
	return
}
