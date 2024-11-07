package main

import (
	"fmt"
	generics "genericList/Generics"
)

// لیست را می بریم به یک پوشه دیگر
// type List[T any] struct {
// 	Items []T
// }

// type Person struct {
// 	Name string
// 	Age int
// }

func main() {
	// حتما باید تابع ها را اینجا کال کنیم تا اجرا شود
	genericInt()
	genericString()

}



func genericInt(){
	list1 := generics.List[int]{Items: []int{}}
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	list1.Add(4)

	fmt.Println(list1.Items)

	list1.InsertAt(2 , 99)
	fmt.Println(list1.Items)

	list1.Remove(2)
	fmt.Println(list1.Items)

	fmt.Printf("Item index : %d\n " , list1.Find(2))

	list1.RemoveByItem(4)
	fmt.Println(list1.Items)
}



func genericString(){
	list2 := generics.List[string]{Items:[]string{}}
	list2.Add("Ali")
	list2.Add("vahid")
	list2.Add("hasan")
	list2.Add("mehrdad")
	list2.Add("omid")


	fmt.Println(list2.Items)

	list2.Remove(1)
	fmt.Println(list2.Items)

	list2.InsertAt(0 , "Reza")
	fmt.Println(list2.Items)

	list2.RemoveByItem("Ali")
	fmt.Println(list2.Items)



}


// تمامی متد هایی که نوشتیم را به پوشه جنرک می بریم

// func (list *List[T]) Add(item T){ // *List[T] این لیست من باید جنریک باشد
// 	list.Items = append(list.Items, item)
// }


// // یک ایتم را در یک ایندسک مشخص وارد می کنیم
// func (list *List[T]) InsertAt( index int ,item T){ 
// 	// 1 2 3 4 5 6 7 8
// 	//1 2 3 99 4 5 6 7 8
// 	// اولین کار یک واحد به ظرفیت اسلایس خود اضافه می کنیم 
// 	list.Items = append(list.Items,item) // 1 2 3 4 5 6 7 8 99
// 	copy(list.Items[index + 1 : ],list.Items[index :]) // 1,2,3,3,4,5,6,7,8
// 	list.Items[index] = item
// }

// // هم با کپی می توانیم انجام بدیم هم با اپند 
// // حذف براساس ایندکس
// func (list *List[T]) Remove(index int){
// 	//copy(list.Items[index : ] , list.Items[index +1 :])
// 	list.Items = append(list.Items[:index], list.Items[index + 1:]...)
	
// }

// func (list *List[T]) RemoveByItem(item T){
// 	index := list.Find(item)
// 	if index != -1 {
// 		list.Remove(index)

// 	}
// }

// //go get -u github.com/google/go-cmp ادرس لایبرری کامپ برای مقایسه 
// func (list *List[T]) Find(item T) int{
// 	for i , v := range list.Items{

// 		// if v == item{  // در جنریک برای بعضی از عملگرهای مقایسه ی خطا می دهد  برای حل این مشکل می ایم از لایبرری کامپ برای مقایسه استفاده می کنیم
// 		// 	return i
// 		// }
// 		if cmp.Equal(v , item){
// 			return i
// 		} 
// 	}
// 	return -1
// }
