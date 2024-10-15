package main

import (					
	"fmt"
	"strings"
)

func main() {
	myString := "this is golang tutorial go go"
	fmt.Println(strings.Contains(myString,"go")) // برای جستجوی یک عبارت داخل یک رشته ،جواب ترو یا فالس است
	fmt.Println(strings.Contains(myString,"go11"))
	fmt.Println(strings.ContainsAny(myString,"go1")) //اگر فقط یکی از این کلمات داخل عبارت بود ترو برمی گرداند 
	fmt.Println(strings.ContainsAny(myString , "1")) 

	fmt.Println(strings.Count(myString , "go")) // تعداد گو ها را می شمارد 

	fmt.Println(strings.Cut(myString , "go")) // اولین گو را که پیدا کرد پاک میکند
	fmt.Println(strings.Cut(myString , "go1")) // اولین گو را که پیدا کرد پاک میکند

	mystringArray := strings.Split(myString , " ")

	println("word count: ",len(mystringArray)) // تعداد کلمات ارایه را برای ما می شمارد

	for _,word := range mystringArray {
		println(word)
	}

	mystringArray2 := strings.Join(mystringArray,"-")

	println(mystringArray2)


	println(strings.Repeat("vahid ", 10))	

	println(strings.Replace(myString , "go" , "golang" ,2)) // کلمه گو را تا دوبار در متن با گولنگ عوض می کند 
	println(strings.ReplaceAll(myString , "go","golang"))


	println(strings.Compare("golang", "golang")) //خروجی 0 یعنی برابر
	println(strings.Compare("Golang", "golang"))//خروجی -1 یعنی سمت چپ بزرگتر
	println(strings.Compare("golang", "GOLANG"))// خروجی 1 یعنی سمت راست بزرگتر

	println(strings.EqualFold("golang", "GOLANG")) // کیس سنستیو نیست

	println(strings.HasPrefix("Iran", "Ir")) // پسوند ایا با ای ار  شروع شده یا نه به حروف بزرگ و کوچک حساس است
	println(strings.HasPrefix("Iran", "ir"))

	println(strings.HasSuffix("Iran", "an")) //پسوند ایا با ای ار  تمام شده یا نه به حروف بزرگ و کوچک حساس است
	println(strings.HasSuffix("Iran", "aN"))

	println(strings.Index("vahid" , "h")) // شماره ایندکس را بر میگرداند


	println(strings.ToLower("Vahid"))
	println(strings.ToUpper("Vahid"))
	println(strings.Title("Vahid lotfi khazineh jadid"))
	

	println(strings.Trim("     Vahid   lotfi" , " "))  // هر چی اسپس دارد حذف می کند
	println(strings.TrimRight("Vahid   lotfi!!" , "!"))
	println(strings.TrimLeft("!!Vahid   lotfi" , "!"))
}

