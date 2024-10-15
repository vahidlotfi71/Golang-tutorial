package main 

func main() {

	defer println("Bye")
	defer println("Goodbye")

	for i := 0 ; i < 10 ; i++ {
		defer println(i)   // برعکس چاپ می کند
	}

	println("hello")
}