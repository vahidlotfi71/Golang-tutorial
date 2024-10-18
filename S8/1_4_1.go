package main

import "fmt"

func main() {
	PrintLog("VAhid", 1, 2, "Hasan", 4, 5, false, 8, 9)
}

func PrintLog(loges ...interface{}) {

	for index, item := range loges {
		println("index ", index, ":", item) // خروجی این قابل فهم نیست
	} 

	println("################################################################")

	for index, item := range loges {
		fmt.Printf("index: %d, item: %v\n", index, item)
	}

}