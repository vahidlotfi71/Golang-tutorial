package main

import "fmt"

func main() {
	var salery float64
	var minSalery float64 = 7000000
	var taxPercentage float64 = 0

	fmt.Print("Enter Your salery: ")
	fmt.Scan(&salery)

	switch {
		case salery < minSalery:
			taxPercentage = 0 
		case salery > minSalery &&  salery < minSalery *2 :
			taxPercentage = .5
		case salery > minSalery * 3 && salery < salery * 4 :
			taxPercentage = .10
		case salery > minSalery * 4 && salery < salery * 6 :
			taxPercentage = .30
		default :
			taxPercentage = .40
		}
	fmt.Printf("your salery is %f\n" ,salery)		
	fmt.Printf("your tax percentage is %f\n" , taxPercentage)
	fmt.Printf("your salery after tax is %f\n" ,salery -(salery * taxPercentage) )

}