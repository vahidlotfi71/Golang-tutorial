package main

var global = 123

func main() {

	global = 1000 
	{
		var local = 10
		println(local)
	}
	//print(x) در اینجا خطا می دهد چون داخل بلاک نیست
	println(global)

	{
		global = global + 10 
		println(global)
	}
}

