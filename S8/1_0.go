package main

func main() {
	Print1()
	str :=  Print2() // چون خروجی دارد به این صورت خروجیش را چاپ می کنیم
	println(str)
	println(Print2())
	println(Print3("Hello world 3"))
	println(Print4("33"))


}

// without input and output
func Print1(){
	println("Hello world1")
}

// without input and output
func Print2() string {
	return "Hello world2"
}

// with input and output
func Print3(Enter string) string {
	return Enter
}

// with input and output
func Print4(Enter string) string {
	return "Hello world" +" " +Enter
} 