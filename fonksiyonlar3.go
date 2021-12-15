package main

import "fmt"

func main() {
	fmt.Println(topla(3, 5))
	fmt.Println(hepsinitopla(3, 65, 2, 74, 12))

}
func topla(x, y int) int {
	return x + y
}
func hepsinitopla(sayilar ...int) int {
	toplam := 0
	for i := range sayilar {
		toplam += sayilar[i]
	}
	return toplam
}
