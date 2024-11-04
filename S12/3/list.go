package main

import "fmt"

// تایپ من جنریک است و استراک است و درون خودش یک اسلایس دارد از جنس تی
type List[T any] struct {
	Items []T
}

func main() {
	list1 := List[int]{Items: []int{}} // چون در بالا گفتم جنس لیست انی است  اینجا می توانیم بگیم جنس لییست اینت است و مقدار ایتم ها را هم جنسش را مشخص کنیم

	list1.Add(1)
	list1.Add(2)
	list1.Add(3)

	fmt.Println(list1.Items)

	list1.InsertAt(1 , 200)
	fmt.Println(list1.Items)

	list1.RemoveAt(1)
	fmt.Println(list1.Items)

}

// می خواهیم اد کنیم به لیست خودمان
// رسیور لیست است 
func (list *List[T]) Add(item T){
	list.Items = append(list.Items  , item)
}

func (list *List[T]) InsertAt(index int , item T  ){
	// 1 2 3 4 5
	// 1 2 99 4 5
	list.Items = append(list.Items, item)				 // 1 2 3 4 5 99 یک واحد به ظرفیتش اضافه کریم
	copy(list.Items[index+1 :] , list.Items[index : ])  // اعداد بعد ایندکس را یک خانه به جلو می بریم
	list.Items[index] = item 							//عدد مورد نظر را در ایندکس مورد نظر می گذاریم

}


func (list *List[T]) RemoveAt(index int) {
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
}
