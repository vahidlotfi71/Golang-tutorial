package main

import (
    "fmt"
    "myproject/math" // مسیر کامل پکیج utils
    
)

func main() {
    result := math.Add(3, 7) // فراخوانی تابع Add از پکیج utils
    result1 := math.Div(4 ,2)
    fmt.Println("Result:", result,"Result1:", result1)
}
