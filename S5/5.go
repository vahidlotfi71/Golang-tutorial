package main

import "fmt"

func main() {
	var salery float64
	var minSalery float64 = 5600000
	var taxPersent float64 = 0


	fmt.Print("Enter Salery: ")
	fmt.Scanln(&salery)

	switch {
	case salery <= minSalery:
		taxPersent = 0 
	case salery > minSalery && salery <= minSalery * 2:
		taxPersent = 0.05
	case salery < minSalery * 2 && salery <= minSalery * 3:
		taxPersent = 0.09
	default:
		taxPersent = 0.15
	}

	fmt.Printf("Your taxPercent is: %f \n",taxPersent)
	fmt.Printf("Your tax is : %.2f\n",salery * taxPersent) //  %.2f\ = تعداد رقم های اعشار فلوت را میدهیم
	fmt.Printf("Your salery is : %.2f\n",salery - taxPersent * salery)
}