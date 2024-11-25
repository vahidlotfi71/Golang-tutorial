package main

import "fmt"

type Room1 struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Status   bool   `json:"status"`
	BedCount int    `json:"bedCount"`
	Price    int    `json:"price"`
}

func main() {

	rooms := []Room1{}
	rooms = append(rooms, Room1{ID: 12 , Type: "standard", Status:false , BedCount: 2 , Price:20000000})


	fmt.Printf("%v\n", rooms.P)

}