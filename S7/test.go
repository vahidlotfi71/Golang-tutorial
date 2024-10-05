package main

import "fmt"

func main() {
	names := [4]string{"a", "b", "c", "d"}

	searchKeyword := "d"

	for index, value := range names {
		if searchKeyword == value {
			fmt.Printf("name found and index is : %v\n", index)
		}
	}

	searchKeyword2 := "a"

	for index , value := range names {
		if searchKeyword2 == value {
			fmt.Printf("name found and index is : %v", index)
		}}
}