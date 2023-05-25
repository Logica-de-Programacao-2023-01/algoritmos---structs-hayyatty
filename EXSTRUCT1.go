package main

import (
	"fmt"
	"math"
)

type circulo struct {
	area float64
	d    float64
}

func calcular(medidas circulo) float64 {
	pi := math.Pi
	raio := medidas.d / 2
	area := pi * (raio * raio)
	return area
}

func main() {
	var (
		d float64
	)
	fmt.Println("Qual o diâmetro do circulo ?")
	fmt.Scan(&d)
	circulo2 := circulo{0, d}
	areacirc := calcular(circulo2)
	fmt.Println("A área do ciruclo de diametro", d, ",é :", areacirc)
}
