// crud

package main 

import "fmt"

type Person struct {
	name string
	family string
	age int 
}

 
// می خواهم اطلات پرسن را در مپ نگه داری کنیم
func main(){

	persons :=make(map[string]Person)

	persons["2960220481"] = Person{name:"vahid" , family:"lotfy" , age:32}
	persons["2960229988"] = Person{name:"hasan" , family:"mehtari" , age : 28}
	persons["5074598745"] = Person{name :"maman" , family:"mosvi" , age: 57}

	fmt.Printf("%v\n",persons)

	persons["5074598745"] =  Person{name :"Ali" , family:"mokarian" , age:2}  // آپدیت یک ردیف
	
	fmt.Printf("%v\n",persons)

	delete(persons, "2960229988") // تابع دیلیت مپ و کلید من را می گیرد و اون ردیف راحذف می کند
	
	fmt.Printf("%v \n" , persons)
	
	

	vahid := persons["2960220481"] // می خواهیم چک کنیم ببینم این کد ملی هست یا نه
	mohammad := persons["5078899655"] // چون همچین چیزی وجود ندارد 0 بر می گرداند

	fmt.Printf("%v \n" , vahid)
	fmt.Printf("%v \n" , mohammad)
 
	mohammad  , ok := persons["5078899655"]
	if ok{
		fmt.Printf("%v \n" , mohammad)
	} else{
		fmt.Printf("not found")
	}




}