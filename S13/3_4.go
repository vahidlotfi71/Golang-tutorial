package main

import (
	"fmt"
	"runtime/debug"
	// "log"	
)

// import "fmt"

func main() {
	defer func(){
		fmt.Println("Start of defer main")

		if err := recover(); err != nil {
			fmt.Println("error: ", err)
			fmt.Println("Recovering from error....")
			debug.PrintStack()
			// debug.Stack() // محتوای استک تریس را به صورت یک بایت اری برمی گرداند
		}
	}()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	function1(nums , 10)     // چون ایندکس خارج از بازه اسلایس می باشد پنیک تولید می کند

	divide(12 , 2)


	// اینجا پنیک اتفاق می افتند چون مخرج صفر است
	divide(32 , 0)


}



func function1(numbers []int, index int) {
	defer func() {
		fmt.Println("Start of Defer function1")
	}()
	// index out of bounds
	// fmt.Println(numbers[index])
	// I create a panic message
	panic("test panic ")  // ایجاد پنیک به صورت دستی
	// I create log Fatalf
	// log.Fatal("test fatal")
	// fmt.Println(12)

}  


func divide( a , b int) /*int*/{
	// این دیفر را به جای اینکه در اینجا بنویسیم می توانیم در در خود تابع مین بنویسیم 
	// defer func(){
	// 	fmt.Printf("defer of divide %d , %d\n" , a , b)

	// 	if err := recover(); err != nil {
	// 		fmt.Println("error: ", err)
	// 		fmt.Println("Recovering from error....")
	// 		debug.PrintStack()
	// 		// debug.Stack() // محتوای استک تریس را به صورت یک بایت اری برمی گرداند

	// 	}
	// }() // همینجا کال می کنیم

	// ایننجا فقط این تابع دیفر را می نویسم

	defer func() {
		fmt.Println("Defer of divide")
	}()

	fmt.Printf("start of divide : %d , %d \n", a, b)
	fmt.Printf("result %d : \n" ,a / b)
	fmt.Printf("End of divide : %d , %d \n\n", a, b)

}
