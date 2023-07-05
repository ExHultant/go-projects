package main

import (
	"fmt"
)

func main() {
	x := 50
	y := 10
	resul := x + y
	resul = x - y
	resul = x % y
	fmt.Print("El Resultado de Y y X es: ", resul)
	fmt.Print("El Resultado de Y y X es: ", resul)
	fmt.Print("El Resultado de Y y X es: ", resul)
	x++
	fmt.Print("El Resultado de X es: ", x)
	x--
	fmt.Print("El Resultado de x es: ", x)
	const pi = 3.14159
	var z = "Mensualidad"
	var w = 80
	//Mostrando con printf
	fmt.Print("%s es %d$\n", z, w)

}
