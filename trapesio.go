package main

import (
	"fmt"
	"math"
)

func trapesio() {
	var option int
	fmt.Println("Seleccione una opcion: ")
	fmt.Println("1.Calcular el area de un Trapecio")
	fmt.Println("2.Calcular el volumen de una esfera")
	fmt.Scan(&option)
	switch option {
	case 1:
		area := Case1()
		fmt.Println("el area de un trapecio es: ", area)
	case 2:
		volumen := Case2()
		fmt.Println("el volumen de una esfera es: ", volumen)
	}
}
func Case1() float64 {
	var baseMayor, baseMenor, altura float64
	fmt.Printf("Ingrese la longitud de la base mayor del trapecio: ")
	fmt.Scan(&baseMayor)
	fmt.Printf("Ingrese la longitud de la base menor del trapecio: ")
	fmt.Scan(&baseMenor)
	fmt.Printf("Ingrese la altura del trapecio: ")
	fmt.Scan(&altura)
	area := ((baseMayor + baseMenor) * altura) / 2
	return area
}

func Case2() float64 {
	var radio float64
	fmt.Printf("Ingrese el radio de la esfera: ")
	fmt.Scan(radio)
	volumen := (4.0 / 3.0) * math.Pi * math.Pow(radio, 3)
	return volumen
}

func main() {
	trapesio()
}
