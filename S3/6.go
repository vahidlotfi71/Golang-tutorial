package main

import "fmt"

 

// func main1() {
	
// 	println("main1 main1 main1")

// 	i , j := 11,22

// 	var IP *int = &i
// 	var jp *int = &j  // مقدار دهی داینامیک 
// 	i3 := i	     	// مقدار دهی استاتیک 
	
// 	println(i)
// 	println(IP)
// 	println(*IP) // برای جاپ مقداز متغییر از ستاره استفاده می شود
// 	println(j)
// 	println(jp)
// 	println(*jp)	
// 	println(&i) // ادرس خانه متغییر ای
// 	println("================================================================")
// 	println(i3)
// }

func main() {
	i, j := 20, 40
	var IP *int = &i  // اشاره به ادرس متغیر ای دارد 
	var JP *int = &j

	fmt.Println(IP)
	fmt.Println("Adress of i2 :" ,&i)

	fmt.Println(JP)
	fmt.Println(&j)

	i2 := i
	i2 = i2 + 2

	fmt.Println("value of i :" ,i)
	fmt.Println("value of i2: ",i2)

	fmt.Println("Adress of i2 :" , &i2)

	var IP2 *int = &i
	fmt.Println("IP2 :",IP2)

	*IP2 = *IP2 + 2
	fmt.Println("value of IP2:" , &i)

	// *IP2 = اشاره به مقدار ان ادارد نه ادرس
}