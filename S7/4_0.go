package main

import "fmt" 

func main() {
	numbers := []int{1, 2, 3, 4, 5} // numbers یک اسلایس است

	changeNumbers(numbers)

	addItem(numbers)  // عدد 6 به اسلایس اضافه نمی شود ما وقتی یک اسلایس را مس فرستیم داخل یک فانکشن وقتی چیزی بهش اپند می کنیم بیرون دیگر اون را نداریم چرا اپند نمی شود چون وقتی ما وقتی چیزی اپند می کنیم اتفاقی می افتند این است که پشت صحه یک ارایه جدید ساخته می شود را حل این مسعله استفاده از اشاره گر است

	addItem1(&numbers)
	
	fmt.Printf("%v\n " , numbers)
}



// نکته : چون این یک اسلایس است تابع روی اسلایس اعمال مشود ولی اگر ارایه بود اعمال نمی شد اگر داخل[] عدد بنویسیم تبدیل به ارایه ممی شود 

func changeNumbers(numbers []int) {
	for i, _ := range numbers {  // چون اینجا با ایتم کاری نداریم ایگنور می کنیم
		numbers[i] = numbers[i] * 3 
	}
}




func addItem(numbers []int){ // تابع یک اسلایس را می گیرد
	numbers = append(numbers,30 )
}



func addItem1(numbers *[]int){
	*numbers = append(*numbers , 22)
}