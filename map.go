package main

import "fmt"

func main(){
	// var person map[string]string = map[string]string{}
	// 	person ["nama"] = "Pablo"
	// 	person ["alamat"] = "Indonesia"


	person :=  map [string]string{
		"nama" : "Pablo",
		"alamat" : "Indonesia",
	}

	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person)

	book:= make(map[string]string)
	book["title"] = "Naruto"
	book ["Penulis"] = "Masashi kisimoto"
	book ["Terbit"] = "2003"

	fmt.Println(book)
}
