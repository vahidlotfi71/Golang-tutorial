package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // با این تابع تعداد سی پی یو هایی که میخواهیم ازش استفاده کنیم را ست م یکنیم  می تونیم دستی بدیم یا از سیستم برمی داریم

	value := 0

	// وقتی اول فانکشن هامون گو می نویسیم دیگر تابع مین ما چیزی چاپ نمی کند چون اجرای این توابع را گوروتین دیگری به عهده گرفته
	go task1() // وقتی اول این یک گو می نویسیم اجرای این در اختیار یک گوروتین دیگر قرار می گیرد و مستقل اجرا می شود و دیگر کد ما خط به خط اجرا نمی شود
	go task2()
	go task3()
	go task4()

	go func() { //  گوی ال این تابع باعث می شود که مقدار ولیو صفر شود یا یه مقداری بهش اضافه شود و نتیجه متفاوت است چون ولیو شیر است
		value++
	}()

	go func() {
		value += 3
	}()

	go func() {
		value += 5
	}()

	fmt.Println(value)

	time.Sleep(time.Second) // این باعث می شود تابع ما خروجی داششته باشد
}

func task1() {
	fmt.Println("task1")
}

func task2() {
	fmt.Println("task2")
}

func task3() {
	fmt.Println("task3")
}

func task4() {
	fmt.Println("task4")
}
