package main

import "time"

func main() {

	firstname := "ahmad"
	names := []string{"vahid", "hasan", "fatemeh", "pejman","vahid1", "hasan1", "fatemeh1", "pejman1"}

	printfirstnamefunc := func() {
		println(firstname)
	}

	firstname = "Reza"
	printfirstnamefunc() // تابع اینجا اینوک (کال) میشه

	//----------------------------------------------------------------
	for i, item := range names {
		go func(index int , name string ) { //  باعث اجراشدن به صرت هم روند می شود دیگر تابع مین منتظر اجرای این فانکشن نمی ماند
			// و به محض تمام شدن تابع مین دیگر این تابع اجرا نمیشود گو باعت این عمل شده
			println(i, name)
		}(i , item)

	}
	time.Sleep(time.Second * 1)
}