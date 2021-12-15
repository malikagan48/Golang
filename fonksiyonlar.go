package main

import "fmt"

// parametre almayan fonksiyon
func main() {
	fmt.Println("Merhaba ....................")
	kimliksor()
	fmt.Println("güle güle ....................")

}
func kimliksor() {
	fmt.Println("isminiz :")
	var isim string
	fmt.Scanln(&isim)
	fmt.Println("yaşınız : ")
	var yas int
	fmt.Scanln(&yas)
	if yas < 18 {
		fark := 18 - yas
		fmt.Printf("üzgünüz %s yaşınız 18 den küçük olduğu için alamıyoruz. %d sene sonra bekleriz ", isim, fark)

	} else {
		fmt.Printf("Hoşgeldin %s iyi eğlenceler ", isim)
	}
}
