package main

import "fmt"

func main() {
	const pi float64 = 3.14

	// pi = 3.15  نمی توانیم جیز دیگری بهش اساین کنیم خطا می دهد

	const (
		name = "vahid"
		number = 21345
		city = "Tehan"
	)

	const googleBaseUrl = "https://www.google.com"
	const googleMapUrl = "/maps"

	println(pi)
	fmt.Printf("%f\n", pi)

	fmt.Printf("name: %s number: %d city: %s pi: %f \n", name, number, city , pi)



	println(googleBaseUrl , googleMapUrl)
}