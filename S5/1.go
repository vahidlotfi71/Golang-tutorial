package main

import "fmt"

func main() {
	var salery float64
	var minSalery float64 = 5600000
	var taxPersent float64 = 0
	var knowlegebase bool = false

	fmt.Println("Enter Salery: ")
	fmt.Scanln(&salery)

	if salery <= minSalery || knowlegebase  {  // || = or
			taxPersent =  0
	} else if salery > minSalery && taxPersent < minSalery *2 {    // && = and
		taxPersent = .05
	}else if salery > minSalery * 2 && salery < minSalery * 3 {
		taxPersent = 0.07
	}else if salery < minSalery*3 && salery > minSalery *4 {
		taxPersent = 0.10
	}else {
		taxPersent = 0.15
	}

	fmt.Printf("Your taxPercent is: %f \n",taxPersent)
	fmt.Printf("Your tax is : %.2f\n",salery * taxPersent) //  %.2f\ = تعداد رقم های اعشار فلوت را میدهیم
	fmt.Printf("Your salery is : %.2f\n",salery - taxPersent * salery)
}