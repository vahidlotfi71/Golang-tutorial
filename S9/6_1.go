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


	var employees []Employee = []Employee{}	

	for _, employee := range fullTimeEmployees { 
		employees = append(employees, employee)
	}

	for _, employee := range partTimeEmployees {  
		employees = append(employees, employee)
	}

	for _, employee := range shiftEmployees { 	 
		employees = append(employees, employee)
	}

	for _ , employee := range employees {
		salary , tax := employee.SalaryCalculator()
		fmt.Printf("\n\nEmployee (%T) : %v\nSalary : %.2f\nTax : %.2f\n", employee, employee , salary, tax)
	
}}

// declare interface
// اینجا ابسترکشن اضافه شده
type Employee interface {
	SalaryCalculator() (salary float64, tax float64) //hours float64 ساعت را حذف کردیم
}

// سه تا استراکت با توجه به نوع کارمنده ایجا می کنیم
type FullTimeEmployee struct {
	Id           int
	NationalCode string
	FullName     string
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
func (employee FullTimeEmployee) SalaryCalculator() (salary float64, tax float64) {//hours float64 ساعت را حذف کردیم
	salary = BaseSalary + (ExtraHourRate*employee.ExtraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

// قوانیینی که در  انترفیس تعریف کردیم را اینجا رعایت می کنیم
func (employee PartTimeEmployee) SalaryCalculator() (salary float64, tax float64) {//hours float64 ساعت را حذف کردیم
	salary = HourlySalaryRate * employee.PartTimeHours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculator() (salary float64, tax float64) {//hours float64 ساعت را حذف کردیم
	salary = ShiftSalaryRate * employee.ShiftHours
	tax = salary * TaxRate
	salary += tax
	return
}