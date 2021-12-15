package main

import (
	"fmt"
)

func main() {

	hesap := 1000
	for true {
		fmt.Println(`**********İŞLEMLER**********
		1.Para çekme
		2.Para yatırma
		3.Bakiye sorgulama
		4.Çıkış
		
		Lütfen işlem tipini seçiniz:`)
		var islem int
		fmt.Scanln(&islem)
		if islem == 1 {
			fmt.Println("Çekilecek tutarı giriniz :")
			var CekilenMiktar int
			fmt.Scanln(&CekilenMiktar)
			if CekilenMiktar <= hesap {
				fmt.Println("Paranızı kutudan alınız.")
				hesap -= CekilenMiktar
				fmt.Println("Bakiyeniz: ", hesap)
			} else {
				fmt.Println("Bakiyeniz yetersiz")
			}
		} else if islem == 2 {
			fmt.Println("yatırmak istediğiniz miktarı giriniz: ")
			var YatirilanMiktar int
			fmt.Scanln(&YatirilanMiktar)
			hesap += YatirilanMiktar
			fmt.Println("Bakiyeniz: ", hesap)

		} else if islem == 3 {
			fmt.Println("Bakiyeniz: ", hesap)
		} else if islem == 4 {
			fmt.Println("Tekrar bekleriz , İyi günler dileriz")
			break
		} else {
			fmt.Println("Lütfen uygun bir işlem türü seçiniz")
		}

	}
}
