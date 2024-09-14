package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	// بین عنصر 7 و 8 می خواهم 75 را بگذارم
	// 1,2,3,4,5,6,7,8,9,10,11,12,13,14
	// 1,2,3,4,5,6,7,75,8,9,10,11,12,13,14 تمام اعداد بعد 75 یک خانه شیف پیدا کردن به جلو

	numbers = append(numbers, 0) // یک صفر به اخر اضافه می کنیم وقتی شیف پیدا کردن جلو فضا داشته باشند

	_ = copy(numbers[7:], numbers[6:]) // نامبرز 7 به بعد را برز در نامبرز 8 به بعد

	numbers[7] = 75

	fmt.Printf("%v \n ",numbers)

	// می خواهیم 75 را حذف کنم 

	numbers = append(numbers[:7] , numbers[8:]...)
	fmt.Printf("%v \n ",numbers)

	numbers = numbers[:0]  // خالی کردن اسلایس
	fmt.Printf("%v \n" , numbers)
	fmt.Printf("len : %v\n", len(numbers))
	fmt.Printf("capacity : %v\n", cap(numbers)) // کپ اسلایس هنوز وجود دارد 

	numbers = nil // خالی کردن اسلایس ، در این حالت دیگر هیچ فضایی در حافظه برای ان در نظر گرفته نمی شود
	fmt.Printf("%v \n",numbers)
	fmt.Printf("len : %v\n", len(numbers))
	fmt.Printf("capacity : %v\n", cap(numbers)) 





}