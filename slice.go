package main
 
import "fmt"

func main (){ 
		names := [...]string{"Mochamad", "Fadli", "Rudiansyah", "Ganteng", "banget", "aamiinn"}
		
		slice1 := names[4:6]
		fmt.Println(slice1)

		slice2 := names[1:5]
		fmt.Println(slice2)

		slice3 := names[:]
		fmt.Println(slice3)

		slice4 := names[:3]
		fmt.Println(slice4)

		slice5 := names[3:]
		fmt.Println(slice5)

		days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
		dayslice6 := days[:]
		fmt.Println(dayslice6)

		dayslice7 := days[5:]
		fmt.Println(dayslice7)

		dayslice7[0] = "Sabtu baru"
		dayslice7[1] = "Taminglay"
		fmt.Println(dayslice7)
		fmt.Println(days)

		daysBaru := append(dayslice7,"Libur Baru")
		daysBaru[0] = "Sabtu Lama"
		fmt.Println(dayslice7)  
		fmt.Println(daysBaru)
		fmt.Println(days)  

	}