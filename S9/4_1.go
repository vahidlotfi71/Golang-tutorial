package main

import "fmt"

type Employee struct {
	ID            int
	Name          string
	Type          string
	SalaryOfMonth int
}

func main() {
	var employ_1 = Employee{ID: 1, Name: "vahid", Type: "master", SalaryOfMonth: 1000}
	var employ_2 = Employee{ID: 1, Name: "hasan", Type: "master", SalaryOfMonth: 2550}


	salery1 := calcYearSalary(employ_1)
	salery2 := employ_1.calcYearSalary() // method
	salery3 := employ_2.calcYearSalary() 

	fmt.Println(employ_1)

	fmt.Println(salery1)	
	fmt.Println(salery2)
	fmt.Println(salery3)
	
}


// function
func calcYearSalary(employee Employee) int {
	return employee.SalaryOfMonth * 12
}


// method of Employee struct فقط استراکت کارمند می تواند ازش استفاده کند
func (employee *Employee) calcYearSalary() int {  // این را تبدیل به متد کردیم
	return employee.SalaryOfMonth * 12			//متد من می تواند اشاره گر باشد یا نباشد
}												// اگر اشاره گر باشد من اگر یک پروپرتی را تغییر دهم بیرون حس میشود
												// اگر اشاره گر نباشد یک کپی وارد فانکشن می شود