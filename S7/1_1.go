package main 

func main() {
	names := [6]string{"ali", "mohammmad" , "hasan" , "vahid" , "famemeh","ali"}

	searchKey := "hasan"
	for i , name := range names {
		if(searchKey == name){
			println("Name found. index: ", i) 
			break
		}}
	searchKey2 := "lotfi"
	for i , name := range names {
		if(searchKey2 == name){
			print("Name found. index: ",i)
			break
		}
	}	
	println("Name not found")

	searchKey2 = "ali"
	for j:= 0 ; j < len(names) ; j++ {
		if names[j] == searchKey2 {
			println("Name found. index: ",j)
			
		}	
	}
	
}