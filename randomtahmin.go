package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// rastgele sayi

	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	sayi := rand.Intn(max - min)
	hak := 10
	for hak > 0 {

		fmt.Println("1 ile 100 arasında bir sayı tuttum tahmin et :")
		var tahmin int
		fmt.Scanln(&tahmin)

		if tahmin > sayi {
			fmt.Println("Aşağıdan bir tahmin yap")
			hak--
			if hak == 0 {
				fmt.Println("hakkınız bitti")
			}
		} else if tahmin < sayi {
			fmt.Println("Yukarıya bir tahmin yap")
			hak--
			if hak == 0 {
				fmt.Println("hakkınız bitti")
			}

		} else {
			fmt.Println("bildiniz doğru")
			break
		}
	}
}
