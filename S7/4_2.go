// copy function

package main

import "fmt"

// یک اسلایس رادر یک اسایس دیگر کپی می کنیم

func main() {
	numbers_1 := []int{1, 2, 3, 4, 5}
	numbers_2 := []int{20 ,30 ,40 , 50 , 60	}

	count := copy(numbers_1 , numbers_2) // این تابع ورودی اول مقصد و ورودی دوم منبع است و دو تا ورودی باید هم نوع باشند 
			//  تابع کپی ارایه های مقصد را در مبدا می ریزد و خود اررایه های منبع از بین می رود 
	fmt.Printf("%v \n ",numbers_1)
	fmt.Printf("number of copy :%v \n ",count) // تعداد اریه های کپی شده را نشان می هد 
}


