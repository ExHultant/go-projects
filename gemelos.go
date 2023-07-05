package main

import "fmt"

// verifica si dos números son gemelos
func sonGemelos(a, b int) bool {
	return (esPrimo(a) && esPrimo(b) && abs(a-b) == 2)
}

// verifica si un número es primo
func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Obtiene el valor absoluto de un número
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Función principal
func gemelos() {
	var inicio, fin int

	fmt.Print("Ingrese el inicio del rango: ")
	fmt.Scanln(&inicio)

	fmt.Print("Ingrese el fin del rango: ")
	fmt.Scanln(&fin)

	fmt.Println("Números gemelos en el rango", inicio, "a", fin, ":")

	for i := inicio; i <= fin-2; i++ {
		if sonGemelos(i, i+2) {
			fmt.Println(i, "y", i+2)
		}
	}
}

func main() {
	gemelos()
}
