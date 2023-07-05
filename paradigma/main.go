package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mi primer Programa en Go")
	//declaracion de variables
	//Si se declara una variable y no se usa da error el progama
	var x int
	//Segunda forma de declarar variables
	y := 3
	fmt.Print("Valor de Y ", y)
	var z int = 13
	fmt.Print("Valor de Z ", z)
	//tipos de variables
	var a int
	var b string
	var c float32
	var d bool
	//con los dos puntos se inicializa una variable y sin los dos puntos se asignan
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
	var r = "Mensualidad"
	var w = 80
	//Mostrando con printf
	fmt.Printf("%s es %d$\n", r, w)
}

//go build genera un .exe del programa
