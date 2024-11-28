package main 

type Number1 interface{
	int | float32 | float64
}

func main(){
	slic1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	x := Sum2(slic1)
	y := Sum1(slic1)
	z := Sum3(slic1)

	println(x)
	println(y)
	println(z)

}


func Sum1(slic []int) (sum int){
	for _ ,v :=range slic {
		sum += v
	}
	return
}

func Sum2[T int | float64](cls []T)(sum T){
	for _ , v :=range cls{
		sum += v

	}
	return
}

func Sum3[T Number1](slice []T) (sum T){
	for _ , v :=  range slice{
		sum += v
	}
	return
}