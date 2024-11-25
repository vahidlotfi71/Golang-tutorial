package main

import (
	"sync"
	"time"
)

var userList = []int{}
var ready = false

func main() {
	// condition1 := sync.Cond{L : &sync.Mutex{} }  // یک وروددی هم می گیرد مه خودش درون خودش از یک لاکر استفاده می کند چون این خودش اشاره گر هست دیگر نیاز نیست از امپرسان استفاده کنیم
	condition:= sync.NewCond(&sync.Mutex{}) //  به این صورت هم ی توانیم کاندیشن تعریف کنیم.و بهتر است این روش برای تعریف کاندیشن استفاده کنیم

	for i:= 0 ; i < 1000 ; i++ {
		NewRequest(i , condition)

	}
	time.Sleep(time.Second*5)
}


func NewRequest(userId int  , condition *sync.Cond){
	Checking(userId , condition)
	condition.L.Lock()  // اینجا لاک م یکنیم
	defer condition.L.Unlock() // condition . locker . lock
	if !ready{  // اگه مخالف ردی بود صبر کن 
		condition.Wait()
	}
	println("user Id: " ,userId, "start streaming")


}

func Checking(userId int  , condition *sync.Cond){
	println(userId , "waiting for start streaming")
	time.Sleep(time.Millisecond * 100) // wait for 100 milisec
	condition.L.Lock()
	defer condition.L.Unlock()	
	userList = append(userList, userId) 
	if len(userList) == 55 {
		ready = true
		condition.Broadcast()  // به همه گوروتین هایی که منتظرن بگو که برن شرط رو چک کنن
		println("User ", userId, "waiting end")

	}

}