package main


func main() {

	i := 0
	list := []int{0,1,2,3,4,5,6,7,8,9}

	for {
		println("Hello vahid ")
		break
	}

	println("--------------------------------")

	for i < 10 {
		println("Hello lotfi" , i)
		i ++  // بعلاوه 1 می کند 
	}

	println("--------------------------------")

	for j := 10 ; j >= 0 ; j--{
		println("hello world" , j)
	}

	println("--------------------------------")

	for index , item := range list {
		println("hello vahid" , index , item)
	}
}
