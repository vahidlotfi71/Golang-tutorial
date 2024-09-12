package main

import "fmt"

func main() {

	name := "vahid"
	age := 32
	nationalcade := 2960220481
	score := 6.6

	print("my name: ", name , " and my age: ", age , " and my score: ", score	, " and my nationalcades: ", nationalcade , "\n")
	println("my name:", name , "and my age:", age , "and my score:", score	, "and my nationalcades:", nationalcade) // پرینت ال بعد اینکه اجرا شد می رود به خط بعد و بین متغییر ها یک اسپس می گذارد
	fmt.Printf("my name: %s and my age: %d and my score: %f and my nationalcades: %d \n" , name, age, score , nationalcade)

	fmt.Printf("name: my type is %T \n",name)
	fmt.Printf("nationalcades binnary is: %b :",nationalcade)

	

}