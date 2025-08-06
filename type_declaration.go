package main

import "fmt"

func main(){
	type NoKTP string
	var ktpPablo NoKTP = "12345678890"

	var xample = "0987654321"
  var xampleKTP = NoKTP(xample) 

	fmt.Println(ktpPablo)
	fmt.Println(xampleKTP)
	fmt.Println(xample)

}
