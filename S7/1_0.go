package main

import "fmt"

func main() {
	var myArr0 [4]int 				// مقادیر 0 ،چون بهش مقدار نادایم 
	var myArr1 [3]int = [3]int{1, 2,3} 
	myArr2 := [5]int{1, 2} 			// چون طول ارایه 5 است و ما دو ا مقدار بهش دادیم بقیه مقادیر را خوش صفر می دهد 
	myArr3 := [...]int{1, 2 ,3 ,5}  // طول ارایه را مشخص نکردیم 

	myArr0[2] = 22
	fmt.Println(myArr0)
	fmt.Println("my array 0 len is : ",len(myArr0))
	println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	fmt.Println(myArr1)
	fmt.Println("my array 1 len is : ",len(myArr1))
	println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	myArr2[4] = 11
	fmt.Println(myArr2)
	fmt.Println("my array 2 len is : ",len(myArr2))
	println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println(myArr3)
	fmt.Println("my array 3 len is : ",len(myArr3))

}