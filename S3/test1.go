package main 

func main() {
	i , j := 11,22

	var IP *int = &i
	var jp *int = &j  // مقدار دهی داینامیک 
	i3 := i	     	// مقدار دهی استاتیک 
	
	println(i)
	println(IP)
	println(*IP) // برای جاپ مقداز متغییر از ستاره استفاده می شود
	println(j)
	println(jp)
	println(*jp)	
	println(&i) // ادرس خانه متغییر ای
	println("================================================================")
	println(i3)
}