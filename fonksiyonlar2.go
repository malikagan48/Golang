package main

import "fmt"

func main() {

	fmt.Println(kareAl(25))
	fmt.Println(Topla(12, 25))
	fmt.Println(Bol(10))
	fmt.Println(Cikar(20))
	selamla("kağan")
}
func kareAl(x int) int {
	return x * x
}
func Topla(x, y int) int {
	return x + y
}
func Bol(x int) int {
	return x / x
}
func Cikar(x int) int {
	return x - x
}
func selamla(isim string) {
	fmt.Printf("%s merhaba", isim)
}
