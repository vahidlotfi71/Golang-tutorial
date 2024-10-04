package main 

func main() {
// 	i := 2
// 	for i < 10 {
// 		println(i)
// 		i = i+1
// 	}
// 	println("====``")
// 	for j := 3 ; j <= 10 ; j++ {
// 		println(j)}

for i := range 10 {
	if i%2 != 0{
		continue
	}
	println(i)
}

}
