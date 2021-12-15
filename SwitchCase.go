package main

import "fmt"

func main() {

	fmt.Println("Ne zaman gideceksin")
	var gun int
	fmt.Scanln(&gun)
	switch gun {
	case 1:
		fmt.Println("Pazartesimi")
	case 2:
		fmt.Println("Salımı")
	case 3:
		fmt.Println("Çarşambamı")
	case 4:
		fmt.Println("Perşembemi")
	case 5:
		fmt.Println("Cumamı")
	case 6:
		fmt.Println("Cumartesimi")
	case 7:
		fmt.Println("Pazarmı")
	default:
		fmt.Println("Geçersiz")
	}
}
