package main 

import ("fmt" )

func main() {
	sum , multiply := Calculator(1,2,3,4,5,6,7,8,9,10)
	sum1 , multiply1 := Calculator2("Vahid",1,2,3,4,5,6,7,8,9,10)


	fmt.Printf("sum:%v, multiply:%f\n",sum,multiply) //خروجی ضرب صفر می شود چرا ؟ چون گفتیم در گو اگر یکر متفییر مقدار دهی اولیه نشه مقدارش صفر می شود
	fmt.Printf("sum1:%v, multiply1:%f\n",sum1,multiply1)

}

func Calculator(numbers ...int) (sum int , mul float64 ) {  // جنس نامبر از نوع اسلایس است
	mul = 1 // برای این مقدار 1  می دهیم که موقع ضرب مقدار ضرب برابر صفر نشود
	for _ , number := range numbers {
		sum += number
		mul *= float64(number)
	}
	return 

}

func Calculator2(name string , numbers ...int) (sum int , mul float64 ) {  // می توانیم یه مقدار در ورودی بدهم ولی در خروجی ازش استفاده نکنیم
	mul = 1 // برای این مقدار 1  می دهیم که موقع ضرب مقدار ضرب برابر صفر نشود
	for _ , number := range numbers {
		sum += number
		mul *= float64(number)
	}
	return 

}