package main

import "fmt"

type triangulo struct {
	base   float64
	altura float64
}

func dadostriangulo() triangulo {
	var (
		base, altura float64
	)
	fmt.Println("Digite qual a base do triangulo :")
	fmt.Scan(&base)
	fmt.Println("Digite qual a altura do triangulo :")
	fmt.Scan(&altura)
	novotriangulo := triangulo{
		base: base, altura: altura,
	}
	return novotriangulo
}
func calculo(t triangulo) {
	area := 0.5 * t.base * t.altura
	fmt.Println("A área do triângulo é:", area)
}

func main() {
	t := dadostriangulo()
	calculo(t)

}
