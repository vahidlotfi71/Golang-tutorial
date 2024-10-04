package main

import "fmt"

func main() {
	// map creation

	names := make(map[string]string) // در این مپ کلید ها استرینگ  و ولیوها استرینق هستن
	names_1 := map[string]string{}
	// var names_2 map[string]string // این خطا می دهد چون فضایی برای دیتا در نظر نگرفتیم
	var names_2 map[string]string = map[string]string{}


	// دیتا وارد مپ کردیم
	names["2960220481"] = "vahid lotfi" 
	names_1["2960220481"] = "vahid lotfi"
	names_2["2960220481"] = "vahid lotfi"

	fmt.Printf("%v \n ",names)
	fmt.Printf("%v \n ",names_1)
	fmt.Printf("%v \n ",names_2)


}