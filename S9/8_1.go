package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name   string
	Color  string
	Lenght int
	wight  int
	Weight int
	Price  int
	Brand  string
	MedIn  string
}

// به این روش میگوییم امبدداستراکت یک استراکت را در دل یک استراکت دیگ استفاده می کنیم

type ElectronicProduct struct {
	Product                // embedded struct  حالا به ویژگی های پروداکت دسترسی داریم
	Ram                    int
	Cpu                    int
	ScreenSize             float64
	OperatingSystem        string
	OperatingSystemVersion string
}

type Mobile struct {
	ElectronicProduct // embedded struct  // می توانیم تایپ استراکت را نگذاریم در این صورت به صورت مستقیم به ویژگی های استراکت دسترسی داریم
	SimCardCapecity   int
	SimCardType       string
	NetworkType       string
	CameraCount       int
}

type Laptop struct {
	ElectronicProduct ElectronicProduct // embedded  struct   حالا به پروپرتی های الکترومیک پروداکت دسترسی داریم برای دسترسی ه ویژگی های الکترونیک پروداکت باید به صورت سلسله مراتبی بریم جلو چون بهش فغحث دادیم
	UsbPortCount      int
	KeyboardType      string
	HasCdRom          bool // have CDROM or not
}

func main() {

	phone := Mobile{}
	phone.Name = "iphone 15"
	phone.Brand = "iphone"
	phone.Color = "black"
	phone.Ram = 12
	phone.OperatingSystem = "ios"
	phone.OperatingSystemVersion = "12"
	phone.CameraCount = 3
	phone.NetworkType = "5G"
	phone.SimCardType = "Nano"
	phone.SimCardCapecity = 2

	fmt.Printf("%+v\n", phone) 
	// برای  نمایش بهتر پرینت بالا تیدیل به جیسون می کنیم

	mobileJson , _ := json.Marshal(phone)
	println(string(mobileJson))
}
