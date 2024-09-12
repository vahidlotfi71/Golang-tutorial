package main

func main() {
	 
	// چاپ اعداد زوج 

	for i:= 0 ; i<= 100 ; i++ {
		if i%2 == 0{
			println(i)
		}
	}

	println("====================================")

	// روش دیگر چاپ اعداد زوج

	for i:= 0 ; i<= 100 ; i++{  
		if i%2 != 0{ // اگه زوج نبود برو اول فر
			continue
		}
		println(i)
	}

println("====================================")

	for i := 0 ; i <= 100 ; i++ { // تا 50 چاپ می کند
		if i == 50 {
			break
		}
		println(i)
	}

}