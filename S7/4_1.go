package main

import (
	"fmt"
	"strings"
)

func main() {

	names := []string{"vahid", "hasan", "ali", "reza"}

	for _, item := range names {  // این فر روی اسلایس عمل نمی کند به این نکته توجه شود وقتی ما  وقتی در فر روی اسلایس کار می کنیم و از ایتم استفاده می کنیم این تغییرات اعمال نمی شود باید مستقیم نیمز را تغییر دهیم
		item = strings.ToUpper(item)
	}

	fmt.Printf("%v\n", names)


	for i , item := range names { // توجه شود وقتی از روی اسلایس کار می کنیم برای ایجاد تغییر مستقیم ار خو اسلایس استفاده کنیم
		names[i] = strings.ToUpper(item)  // در حلقه فر یک کپی از اسلایس ما گرفته می شود و تغییر روی اونن کپی اعمال می شود ایتم یک کپی از اون است
	}

	fmt.Printf("%v\n", names)
}