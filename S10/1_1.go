package main

import "fmt"

type IPerson interface {
	print()
}

type Person struct {
}

func main() {
	var P IPerson
	// پرسن ایمپلیمنت نکرده ای پرسن را  برای حل این مشکل یک فانکشن می نویسیم با ریسیور پرسن و اسم متدش
	P = Person{}
	P.print()

	print2(P)
	print2(12)
	print2("Ali")
	print2(12.14)


}

// ایمپلیمنت یک اینتر فیس در گو به صورت ضمنی است یعی بیرون از تابع اصلی تعریف می شود
// کامپایر خودش حواسش است که پرسن ای پرسن را پیداده سازی کرده
func (person Person) print() {
	println("Hello world!")
}

func print2(item interface{}){  // اینترفیس خالی یعنی هر ورودی می توانیم بهش بدیم
	fmt.Printf("item: %v\n",item)  
}
