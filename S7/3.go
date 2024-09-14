package main

import "fmt"

func main() {

	myArray := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// mySlice := []int{1,2,3,4,5}  // طول اسلایس مشخص نیست

	// mySlice1 := make([]int  , 6) // تابعی است که می توانیم چیزهای که می خواهیم را با هاش بسازیم

	// mySlice2 := make([]int , 8 , 16) //  یکی طول و دیگری ظرفیت است

	mySlice3 := myArray[:] // اسلایس من کل ارایه من است
	mySlice4 := myArray[:5]
	myArray[0] = 555

	fmt.Printf("%v\n", myArray)
	fmt.Printf("%v \n", mySlice3)
	fmt.Printf("%v \n", mySlice4)

	fmt.Println("my slice4 len is: " , len(mySlice4))
	fmt.Println("my slice4 cap is: " , cap(mySlice4))

	fmt.Println("my array len is: " , len(myArray))
	fmt.Println("my array cap is: " , cap(myArray))

	println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	mySlice4 = append(mySlice4 , 22)

	fmt.Printf("%v \n", mySlice4)
	fmt.Println("my slice4 len is: " , len(mySlice4))
	fmt.Println("my slice4 cap is: " , cap(mySlice4))

	 


}