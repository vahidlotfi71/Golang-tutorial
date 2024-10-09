package main 

import ("fmt" )

func main () {
	numbers := [4]int{1, 2, 3, 4}
	numbers2 := numbers

	fmt.Println(numbers)    // برای چاپ ارایه از پکیج اف ام تی استفاده شود
	fmt.Println(numbers2)

	println(&numbers) // & برا ی چاپ ادرس خوانهایی که برای این ها رزرو شده استفاده می شود 
	println(&numbers2)

//================================================================
	changevalue(numbers)
	changevalue(numbers2)	

	fmt.Println(numbers)    // برای چاپ ارایه از پکیج اف ائ تی استفاده شود
	fmt.Println(numbers2)
}


func changevalue(myArray [4]int){  // یک کپی از از ارایه گفته شده بنابرین تغییر در اینجا هیچ تغییری در ارایه اصلی ایجاد نمی کند
	myArray[0] = 777 
	myArray[1] =888
}
