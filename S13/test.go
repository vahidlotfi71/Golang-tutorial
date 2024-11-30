package main

import (

	"fmt"
	"io"
	"os"
)

func main() {
	error := Copy("str.txt" , "des.txt")
	
	if error != nil {
		println("Error copying str.txt file to file: ", error)
	}
}

func Copy(source, destination string) error {
	sou, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("an error occurred while opening %s: %v", source, err)
	}

	des , err := os.Create(destination)

	if err != nil {
		return fmt.Errorf("an error occurred while creating %s: %v", destination, err)
	}

	_ , err = io.Copy(des, sou)

	if err != nil {
		return fmt.Errorf("an error occurred while copying %s: %v", destination, err)
	}
	
	defer des.Close()
	defer sou.Close()
	

	return nil 

}