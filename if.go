package main

import "fmt"

func main() {
	name := "Pablo"

	if name == "Pablo" {
		fmt.Println("Hello Pablo")
	} else if name == "Amell" {
		fmt.Println("Hi, Cantik")
	} else if name == "Lia" {
		fmt.Println("Hi, Lia")
	} else if name == "yaya" {
		fmt.Println("Hi, yaya")
	} else {
		fmt.Println("siapa kamu?")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
