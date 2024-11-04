package main 

import "fmt"


type Number interface {
	int | float64
}


func main() {

	mySlic1 := []int{1, 2, 3, 4 , 5, 6,}
	mySlic2 := []float64{1.5 , 2.5, 3.5, 4.5, 5.5}
	mySlic3 := []float64{1 , 2, 3, 4.5, 5.5}


	fmt.Println(Sum1(mySlic1))
	fmt.Println(Sum2(mySlic2))
	fmt.Println(Sum3(mySlic3))


}



// یک اسلایس از ورودی می گیریم و جمع اسلایس را می دهیم به خروجی

// می خواهیم تابع زیر را تبدیل کنیم به جنریک
func Sum1(slc []int) (sum int){
	for _ , v := range slc{
		sum += v
	}
	return 
}



// جنریک تابع بالا
func Sum2[T int | float64](slc []T) (sum T){
	for _ , v := range slc{
		sum += v
	}
	return 
}


func Sum3[T Number](slc []T) (sum T){
	for _ , v := range slc{
		sum += v
	}
	return 
}
