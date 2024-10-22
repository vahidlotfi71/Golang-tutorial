package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name   string
	Color  string
	lenght int
	wight  int
	Weight int
	Price  int
	Brand  string
	MedIn  string
}

type ElectronicProduct struct {
	Product
	Ram                    int
	Cpu                    int
	ScreenSize             float64
	OperatingSystem        string
	OperatingSystemVersion string
}

type Mobile struct {
	ElectronicProduct
	SimCardCapecity int
	SimCardType     string
	NetworkType     string
	CameraCount     int
}

type Laptop struct {
	ElectronicProduct
	UsbPortCount int
	KeyboardType string
	HasCdRom     bool
}

func main() {

	phone := new(Mobile)
	phone.Name = "sumsung 1212"
	phone.Brand = "sumsung"
	phone.Color = "Red"
	phone.Ram = 12
	phone.OperatingSystemVersion = "Android"
	phone.OperatingSystemVersion = "12"
	phone.CameraCount = 3
	phone.NetworkType = "5G"
	phone.SimCardType = "Nano"
	phone.SimCardCapecity = 2

	fmt.Printf("mobile : %+v\n", phone)

	mobileJson , _ := json.Marshal(phone)

	println(string(mobileJson))

}