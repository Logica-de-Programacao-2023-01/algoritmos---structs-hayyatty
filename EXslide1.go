package main

import "fmt"

type Retangulo struct {
	altura  float64
	largura float64
}

func calculararea(x Retangulo) float64 {
	return x.largura * x.altura
}
func main() {
	var (
		z, y float64
	)
	fmt.Println("Digite a altura :")
	fmt.Scan(&z)
	fmt.Println("Digite a largura :")
	fmt.Scan(&y)
	RetanguloDig := Retangulo{z, y}
	area := calculararea(RetanguloDig)
	fmt.Println("A area desse retangulo é de ", area)
}
