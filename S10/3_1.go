package main

import "fmt"

type Runner interface { // رفتار دویدن دار بهتر است اسمش را رانر بگذاریم و در دلش متد ران را داشته باشم
	Run()
}

// رفتار واک را دارد
type Waker interface {
	walk()
}

type Shooter interface {
	shoot()
}

type Player struct {
	Name     string
	Age      int
	Height   int
	Weight   int
	Position string
}

func main() {

	player1 := &Player{
		Name : "John",
		Age   : 13,
		Height : 200 ,
		Weight : 100,
		Position : "Forward",
	}

	var runner Runner = player1
	//var waker Waker = player1  // برای این متغییر خطا می دهد چون متد این واکر را پایین پیاده سازی نکردیم
	var shooter Shooter = player1	

	runner.Run()
	shooter.shoot()

}

func (player *Player) Run() {  // اگر از ستاره استفاده کنیم ایمپایمنت نمی کند
	fmt.Printf("name: %s , position : %s , Player is running\n\n" , player.Name , player.Position)
}

func (player *Player) shoot() {
	fmt.Printf("name: %s , position : %s , Player is shooting\n\n" , player.Name , player.Position)
}