package main

import "fmt"

const (
	BaseSalary       = 500
	ExtraHourRate    = 10
	HourlySalaryRate = 11
	ShiftSalaryRate  = 8
	TaxRate          = 0.09
)

func main() {
	//
	fullTimeEmployees := []FullTimeEmployee{
		{Id: 1, NationalCode: "1234567890", FullName: "Pejman Rezaee", ExtraHours: 40},
		{Id: 2, NationalCode: "4836524125", FullName: "Maryam Hosseini", ExtraHours: 120},
	}

	partTimeEmployees := []PartTimeEmployee{
		{Id: 3, NationalCode: "6563453455", FullName: "Milad Hassani", PartTimeHours: 100},
		{Id: 4, NationalCode: "5435435435", FullName: "Maryam Rezaee", PartTimeHours: 87},
	}

	shiftEmployees := []ShiftEmployee{
		{Id: 5, NationalCode: "3123123213", FullName: "Shahin", ShiftHours: 150},
		{Id: 6, NationalCode: "6363454355", FullName: "Masoud", ShiftHours: 168},
	}

	// 1
	var employee Employee = fullTimeEmployees[0]
	salary, tax := employee.SalaryCalculator(float64(fullTimeEmployees[0].ExtraHours))
	fmt.Printf("Employee: %v\nSalary: %.2f\nTax: %.2f\n", employee, salary, tax)

	// 2
	employee = fullTimeEmployees[1]
	salary, tax = employee.SalaryCalculator(float64(fullTimeEmployees[1].ExtraHours))
	fmt.Printf("Employee: %v\nSalary: %.2f\nTax: %.2f\n", employee, salary, tax)

	// 3
	employee = partTimeEmployees[0]
	salary, tax = employee.SalaryCalculator(float64(partTimeEmployees[0].PartTimeHours))
	fmt.Printf("Employee: %v\nSalary: %.2f\nTax: %.2f\n", employee, salary, tax)

	// 4
	employee = partTimeEmployees[1]
	salary, tax = employee.SalaryCalculator(float64(partTimeEmployees[1].PartTimeHours))
	fmt.Printf("Employee: %v\nSalary: %.2f\nTax: %.2f\n", employee, salary, tax)

	// 3
	employee = shiftEmployees[0]
	salary, tax = employee.SalaryCalculator(float64(shiftEmployees[0].ShiftHours))
	fmt.Printf("Employee: %v\nSalary: %.2f\nTax: %.2f\n", employee, salary, tax)

	// 4
	employee = shiftEmployees[1]
	salary, tax = employee.SalaryCalculator(float64(shiftEmployees[1].ShiftHours))
	fmt.Printf("Employee: %v\nSalary: %.2f\nTax: %.2f\n", employee, salary, tax)

}

// declare interface
// اینجا ابسترکشن اضافه شده
type Employee interface {
	SalaryCalculator(hours float64) (salary float64, tax float64)
}

// سه تا استراکت با توجه به نوع کارمنده ایجا می کنیم
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

type ShiftEmployee struct {
	Id           int
	NationalCode string
	FullName     string
	ShiftHours   float64
}

// قوانیینی که در  انترفیس تعریف کردیم را اینجا رعایت می کنیم
func (employee FullTimeEmployee) SalaryCalculator(extraHours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*extraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

// قوانیینی که در  انترفیس تعریف کردیم را اینجا رعایت می کنیم
func (employee PartTimeEmployee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = HourlySalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = ShiftSalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}