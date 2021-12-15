package main

import (
	"fmt"
)

func main() {

	hak := 3
	for hak > 0 {
		fmt.Println("İsminiz :")
		var isim string
		fmt.Scanln(&isim)
		fmt.Println("Yaşınız :")
		var yas int
		fmt.Scanln(&yas)

		if yas < 18 {
			fark := 18 - yas
			fmt.Printf("Üzgünüm %s yaşın 18 den küçük olduğu için içeri giremezsin. %d sene sonra sonra bekleriz", isim, fark)
			hak--
			fmt.Println("   ")
			if hak == 0 {
				fmt.Println("Başka bir alım kabul edemiyoruz")
				break
			} else {
				fmt.Println("Sırada ki kimliğe bakalım.")

			}

		} else {
			fmt.Printf("Hoşgeldiniz %s  iyi eğlenceler...", isim)
			fmt.Println("  ")
			if hak == 0 {
				fmt.Println("Başka bir alım kabul edemiyoruz")
				break
			} else {
				fmt.Println("Sırada ki kimliğe bakalım.")

			}

		}

	}
}
