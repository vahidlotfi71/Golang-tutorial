// bad example
// در این کد اصول سالید رعایت نشده است 

package main

import "fmt"

// salery base is duler $
const (
	BaseSalery = 1000
	ExtraPerHourRate = 10
	HourlySaleryRate = 11
	ShiftSaleryRate = 8
	TaxRate = 0.9
)

func main() {
		employees := []Employee{
			{ID: 100 , NationalCode: "1234567890" , FullName:"vahid lotif" ,Type: "fullTime" , Hours :12},
			{ID: 200 , NationalCode: "3543535524" , FullName:"ali samadi", Type: "fullTime" , Hours :9},
			{ID: 300 , NationalCode: "2457884452" , FullName:"fatemeh lotfi", Type: "partTime",Hours :150},
			{ID: 400 , NationalCode: "3445668765" , FullName:"Reza ahmadi", Type: "Shift",Hours :70},

	}

	for _ , employee := range employees{
		salery , tax := employee.saleryCalculetor(float64(employee.Hours))
		fmt.Printf("Employee: %v\n salery: %.2f\n tax: %.2f\n\n", employee , salery , tax)
	}
}


type Employee struct{
	ID int 
	NationalCode string
	FullName string
	Type string
	Hours float64

}


func (employee Employee) FullTimeSaleryCalculator(extraHours float64) (salery float64 ,Tax float64){
	salery = BaseSalery + (ExtraPerHourRate * extraHours) * 1.4
	Tax = salery * TaxRate
	salery += salery + Tax
	return
}


func (employee Employee) PartimeSaleryCalculator(Hours float64) (salery float64 ,Tax float64){
	salery = HourlySaleryRate * Hours
	Tax = salery * TaxRate
	salery += salery + Tax
	return
}

func (employee Employee) ShiftSaleryCalculator(Hours float64) (salery float64 ,Tax float64){
	salery = ShiftSaleryRate * Hours
	Tax = salery * TaxRate
	salery += salery + Tax
	return
}


func (employee Employee) saleryCalculetor(hours float64) (salery float64 ,Tax float64){
	if employee.Type == "fullTime" {
		return employee.FullTimeSaleryCalculator(hours)
	}else if employee.Type == "partTime"{
		return employee.PartimeSaleryCalculator(hours)
	}else {
		return employee.ShiftSaleryCalculator(hours)
	}
}