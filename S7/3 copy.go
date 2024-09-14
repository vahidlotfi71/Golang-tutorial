package main

import "fmt"

func main() {

	myArray := [6]int{1, 2, 3, 4, 5 ,6}

	mySlice4 := myArray[:4]
	myArray[0] = 555

	fmt.Printf("%v\n", myArray)

	fmt.Printf("%v \n", mySlice4)

	fmt.Println("my slice4 len is: " , len(mySlice4))
	fmt.Println("my slice4 cap is: " , cap(mySlice4))

	mySlice4 = append(mySlice4 , 22 ,44)
	mySlice4 = append(mySlice4 , 33)  // ظرفیت دو برابر ظرفیت قبلی می شود

	myArray[1] = 111111  // این تغیر دیگر در اسلایس اعما نمشود 

	fmt.Printf("%v \n", mySlice4)

	fmt.Println("my slice4 len is: " , len(mySlice4))
	fmt.Println("my slice4 cap is: " , cap(mySlice4))

	
	fmt.Printf("%v\n", myArray)
	fmt.Printf("%v \n", mySlice4)

	 


}