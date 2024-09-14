// دوتا اسلایس را مرج می کنیم با استفاده از اپند

package main 

import "fmt"

func main(){
	numbers_1 := []int{1, 2, 3, 4, 5}
	numbers_2 := []int{10 , 20 , 30 , 40 , 50}

	numbers_1 = append(numbers_1, 6,7,8,9)
	fmt.Printf("%v\n", numbers_1)

	numbers_1 = append(numbers_1,numbers_2...) //حتما باید سه نقطه گذاشته شود وگر نه خطا می ده 
	fmt.Printf("%v \n",numbers_1)
}