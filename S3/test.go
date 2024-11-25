package main

func main(){
	i , j := 12,20

	var I = &i
	var J = &j

	i = 100

	println(i, j)
	println(*I)
	println(*J)
}