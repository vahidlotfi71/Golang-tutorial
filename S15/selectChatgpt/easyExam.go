// // مثال ۱: ارسال و دریافت ساده با select (سطح آسان)
// package main

// import (
// 	"fmt"

// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func (){
// 		ch1 <- "سلام از کانال 1"
// 	}()

// 	go func(){
// 		ch2 <- "سلام از کانال 2"
// 	}()

// 	select{
// 	case message1 := <- ch1:
// 		fmt.Println("داده دریافتی از کانال 1: ", message1)
// 	case message2 := <- ch2:
// 		fmt.Println("داده دریافتی از کانال 2 : ", message2)
// 	default:
// 		fmt.Println("ERROR")
// 	}
// }

// استفاده از سلکت با تایمر سطح سوال متوسط

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main(){
// 	ch := make(chan string)

// 	go func () {
// 		time.Sleep(time.Second * 2)
// 			ch <- "داد ارسال شد"
// 	}()

// 	select{
// 	case message := <- ch:
// 		time.Sleep(time.Second * 1)
// 		fmt.Println("داد دریافت شد : ",message)

// 	case <- time.After(time.Second *1):
// 		fmt.Println("زمان تمام شد و هیچ داده ی دریافت نشد")

// 	}

// }

package main

import (
	"fmt"
	"time"
) 

func Worker(id int  , ch chan string){
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Worker id %d : کار تمام شد", id)
}

func main(){
	ch1 := make(chan string) 
	ch2 := make(chan string)
	ch3 := make(chan string)
	
	go Worker(1 , ch1)
	go Worker(2 , ch2)
	go Worker(3 , ch3)	


	for i:= 0; i<3 ; i++ {
		select{
		case message1 := <- ch1:
			fmt.Println(message1)
		case message2 := <- ch2:
			fmt.Println(message2)
		case message3 := <- ch3:
			fmt.Println(message3)

		}

	}

}