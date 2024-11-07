package generics

import "github.com/google/go-cmp/cmp"


type List[T any] struct {
	Items []T
}



func (list *List[T]) Add(item T){ // *List[T] این لیست من باید جنریک باشد
	list.Items = append(list.Items, item)
}


// یک ایتم را در یک ایندسک مشخص وارد می کنیم
func (list *List[T]) InsertAt( index int ,item T){ 
	// 1 2 3 4 5 6 7 8
	//1 2 3 99 4 5 6 7 8
	// اولین کار یک واحد به ظرفیت اسلایس خود اضافه می کنیم 
	list.Items = append(list.Items,item) // 1 2 3 4 5 6 7 8 99
	copy(list.Items[index + 1 : ],list.Items[index :]) // 1,2,3,3,4,5,6,7,8
	list.Items[index] = item
}

// هم با کپی می توانیم انجام بدیم هم با اپند 
// حذف براساس ایندکس
func (list *List[T]) Remove(index int){
	//copy(list.Items[index : ] , list.Items[index +1 :])
	list.Items = append(list.Items[:index], list.Items[index + 1:]...)
	
}

func (list *List[T]) RemoveByItem(item T){
	index := list.Find(item)
	if index != -1 {
		list.Remove(index)

	}
}

//go get -u github.com/google/go-cmp ادرس لایبرری کامپ برای مقایسه 
func (list *List[T]) Find(item T) int{
	for i , v := range list.Items{

		// if v == item{  // در جنریک برای بعضی از عملگرهای مقایسه ی خطا می دهد  برای حل این مشکل می ایم از لایبرری کامپ برای مقایسه استفاده می کنیم
		// 	return i
		// }
		if cmp.Equal(v , item){
			return i
		} 
	}
	return -1
}
