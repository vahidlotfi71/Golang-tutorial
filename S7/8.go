package main 

// جمع کدام اعداد برابر 10 است 
num  := []int{3,4,6,7,6,5,2,3}  
tearget  := []int{10}



// 3,4,16,7,10,15,2,3 اعداد ارایه ما 
// 7,6,-6,3,0, -5,8,7  اعداد بالا با این اعداد جمع بشن 10 می شوند 
// روش حل => ارایه اول را در مپ قرار می دهیم م همچنین مقدار ایندکس ارایه در به عنوان عضور دوم مپ قرار می دهیم
// [3 , 0],[4, 1] ,[6 , 3] => وقتی ما به 6 می رسیم چون قبلا 4  را داشتیم و همچنین مقدار اینکسشم داریم  می فهمییم که 4و 6  ده می شود 

func twoSum (num [] int , target [] int) {
	keys := map[int]int{}

	for i , item := range num{
		complimant := target - num

		if _ ,ok := keys[complimant] ; ok {
			return []int{keys[complimant],i}
			}
			keys[item] = i	
		}
		return []int{0 ,}
	}




}