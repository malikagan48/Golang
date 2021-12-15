package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*x := 5        // işlemlerimize var int değerlerini tanımlamadan da sağlayabiliriz.
	var y int = 4 //int değerlerini döndürmek için var değerini tanımlamalıyız.
	var sum int = x + y
	fmt.Println(sum)

	var isim_soyisim string
	var yas int
	var yetiskinmi bool
	isim_soyisim = "kağan küçük"
	yas = 24
	yetiskinmi = true
	fmt.Println("isim soyisim", isim_soyisim)
	fmt.Println("yas", yas)
	fmt.Println("yetiskin mi", yetiskinmi)
	fmt.Printf("isim soyisim  %s ,yaşımız %d ,yetişkin mii  %t", isim_soyisim, yas, yetiskinmi)

	// for örneği
	toplam := 0
	for i := 0; i < 10; i++ {
		toplam += i
		fmt.Println("toplam :", toplam)
	}
	fmt.Println("toplam :", toplam)
	// while for örneği
	toplam := 1
	for toplam < 1000 {
		toplam += toplam
		fmt.Println("toplam :", toplam)

	// if örneği
	var a int = 9
	if a > 9 {
		fmt.Println("a sayısı 9 dan büyüktür", a)
	} else if a == 9 {
		fmt.Println("a sayısı 9 a eşittir", a)

	} else {
		fmt.Println("a sayısı 9 dan küçüktür", a)
	}


	//input alma örneği
	fmt.Println("isim giriniz")
	var isim string
	fmt.Scanln(&isim)
	fmt.Println("girdiğiniz isim", isim)

	// int değer alma
	fmt.Println("sayı giriniz")
	var sayi int
	fmt.Scanln(&sayi)
	fmt.Println(sayi * sayi)

	*/
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
