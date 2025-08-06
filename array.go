package main

import "fmt"

func main(){
	// var names[3] string
	// names[0] = "Mochamad"
	// names[1] = "Fadli"
  // names [2] = "Rudiansyah"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])
	var names = [...] string{
	 "Mochamad",
	 "Fadli",
   "Rudiansyah",
	}
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
	90,
	88,
	336, 
	225,
	900,
	899,
	999,
}
	fmt.Println(values)
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[0])
	fmt.Println(len(values))
}