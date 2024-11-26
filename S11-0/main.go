package main

import (
	"fmt"
	jalaali "github.com/jalaali/go-jalaali" // حال ماژولی که گت کرده بودیم را به عنوان دیپندنسی های پروژه خود اضافه م یکنم
	"github.com/naeemaei/moduleExampl/services"
	// یک اسم هم برای ماژولی که گت کردیم می گذاریم

	
)

func main() {
	fmt.Println("Hello world!!!")
	year , month ,day , error :=  jalaali.ToGregorian(1371 , 5 ,35 )

	if error == nil {
		fmt.Printf("%d/%d/%d\n" , year , month , day)
	}else{
		fmt.Print(error)
	}

	var services services.TestService = services.TestService{}  // ادرس پکیج داخلی را بهش دادیم
	
	fmt.Printf("%v" , services)





}
