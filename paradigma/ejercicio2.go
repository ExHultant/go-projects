/*
Realizar cualquier programa que por lo menos lea dos variable y haga calculo
*/
package main

import (
	"fmt"
	"strconv"
)

func ejercicio2() {
	var num1, num2 int

	// Lee el primer número
	fmt.Print("Ingrese el primer número: ")
	input1 := getInput()
	num1, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Println("Error al convertir el primer número:", err)
		return
	}

	// Lee el segundo número
	fmt.Print("Ingrese el segundo número: ")
	input2 := getInput()
	num2, err = strconv.Atoi(input2)
	if err != nil {
		fmt.Println("Error al convertir el segundo número:", err)
		return
	}

	// Realiza el cálculo
	resultado := num1 + num2

	// Imprime el resultado convertido a string
	resultadoStr := strconv.Itoa(resultado)
	fmt.Println("El resultado de la suma es:", resultadoStr)
}

// Función auxiliar para leer la entrada del usuario
func getInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}
func main() {
	ejercicio2()
}
