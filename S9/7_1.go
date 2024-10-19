package main

import "fmt"

type Test struct {
	Id   int
	Name string
}

func main() {
	test := Test{Id: 12 , Name: "vahid" }

	test.Print()
	test.print2()
}

func (test Test) Print() { // method Print  //Print  اگر خارج پکیج مین باشیم پرنت را می بینم چون اسم فانکش بل حروف بزرگ اولش نوشته شده است	
	fmt.Println(test.Id , test.Name)
}

func (test Test) print2() { // method print2  //print2  اگر خارج پکیج مین باشیم پرنت را نمی بینم چون اسم فانکش با حروف کوجک اولش نوشته شده است
	fmt.Println(test.Id , test.Name)
}