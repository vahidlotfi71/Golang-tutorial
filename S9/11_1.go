package main

import (
	"encoding/json"
	"encoding/xml"
)

type Person struct {
	Name      string 	`json:"first_name" xml:"NAME"`  // با استراک تگ خروجی جی سون با این تگ هاییاست که دادیم
	Age       int		`json:"age,omitempty" xml:"AGE"`	//omitempty مقدار نداشت در خروجی نمی اورد
	Family    string 	`json:"last_family" xml:"FAMILY"`
	Gender    string    `json:"gender" xml:"GENDER"`
	Height    int		`json:"height" xml:"HEIGHT"`
	Weight    int 		`json:"weight" xml:"WEIGHT"`
	HairColor string 	`json:"-" xml:"HAIRCOLOR"` // با دش گذاشتن دیگر دیتا در خروجی نمایش داده نمشود
}

// می خواهیم پرسون را به فایل جی سون تبدیل کنیم
func main() {
	person := Person{Name: "John", Family: "lotfi", Gender: "Male", Height: 173, Weight: 70, HairColor: "Black"}

	personJson , _ := json.MarshalIndent(person , "", "	")  // اگر پراپرتی هایی که برای تایپ پرسن تعریف کردیم با حروف کوچک شروع شده باشد خروجی این خالی است چون از بیرون بهش دسترسی نداریم


	personxml,_ := xml.MarshalIndent(person , "", "	")

	println(string(personJson))
	println(string(personxml))

}