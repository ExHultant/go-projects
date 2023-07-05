package main

import "fmt"

func main() {
	var sueldo float64
	var impuesto float64
	fmt.Print("Ingrese el numero de Trabajadores: ")
	var numTrabajadores int
	fmt.Scanln(&numTrabajadores)
	for i := 1; i < numTrabajadores; i++ {
		fmt.Printf("Ingrese el sueldo del trabajador %d: ", i)
		fmt.Scanln(&sueldo)
		switch {
		case sueldo < 500:
			impuesto = sueldo * 0.9
		case sueldo > 500:
			impuesto = sueldo * 0.11
		}
		fmt.Printf("El sueldo del trabajador %d es: %.2f\n", i, sueldo)
		fmt.Printf("El impuesto a pagar es: %.2f\n", impuesto)
	}
}
