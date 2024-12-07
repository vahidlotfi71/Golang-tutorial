package main

import (
	"fmt"
)

func main() {

	// ایجاد یک کانال با بافر با ظرفیت 2
	ch := make(chan int, 2)

	// send data to channel
	ch <- 1
	fmt.Println("data 1 sended ")
	ch <- 2
	fmt.Println("data 2 sended ")

	// این خط مسدود می‌شود اگر ظرفیت کانال پر باشد
	go func() {
		ch <- 3
		fmt.Println("data 3 sended")

	}()

	// دریافت داده از کانال

	fmt.Println("data received", <-ch)
	fmt.Println("data received", <-ch)

	// اکنون ظرفیت خالی شده و گوروتین می‌تواند داده 3 را ارسال کند

	fmt.Println("data received", <-ch)

}
