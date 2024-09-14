package main

import "fmt"

func main() {
	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	numbers2 := &numbers

	numbers3 := numbers[:2]

	numbers4 := numbers3

	numbers[0] = 100

	println(&numbers)
	println(&numbers2)
	println(&numbers3)
	println(&numbers4)

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)

	changeValue(&numbers)
	changeValue(numbers2)  // چون نامبر 2 اشاره گر هست دیگر لازم نیست از امپرسان استفاده کنیم 

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)

}

func changeValue(mayArray *[8]int) {    //* از اشاره گر استفاده می کنیم تا به ادرس خانه ارایه اشاره کند و تغییرات ما اعمال شود 
	mayArray[0] = 555
	mayArray[1] = 666
}