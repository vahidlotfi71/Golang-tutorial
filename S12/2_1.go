package main

import "fmt"


type Number interface {
	int | int16 | int32 | int64 |float32 | float64
}

func main() {
	x := SumIns(12, 12)
	y := SumFloats(12.5 , 2.2)
	z := Sum(12 ,12.5)

	d := Div(12 , 2)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)

}

func SumIns(a, b int) int {
	return a + b
}

func SumFloats(a, b float64) float64 {
	return a + b
}

// برای اینکه یک کار مشابه را با انواع مختلف انجام دهیم می اییم جنریک م ینویسیم
		// تی نوعش اینت یا فلوت است و فقط اینت یا فلوت  قبول می کند 
func Sum[T int| float64 ](a , b T) T { //any : به جای دیتاتایپ های مورد قبول تی می توانیم انی هم بگذاریم ولی در عمل جمع خطا می خورد 
	return a + b
}

func Div[T Number ](a , b T) T { 
	return a + b
}

