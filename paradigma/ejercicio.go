 // hallar el area de un triangulo
// hallar el area de un circulo
// hallar el area de un cuadrado
package main

import (
	"fmt"
	"math"
)

func ejercicio() {
	base := 10
	area := base * base
	fmt.Printf("%v", area)
	//Area de un triangulo
	base2 := 10.0 // Inserta aquí el valor de la base del triángulo
	altura := 5.0 // Inserta aquí el valor de la altura del triángulo

	area2 := (base2 * altura) / 2
	fmt.Printf("El área del triángulo es: %.2f\n", area2)
	//Area de un Circulo
	radio := 5.0 // Inserta aquí el valor del radio del círculo

	area3 := math.Pi * math.Pow(radio, 2)
	fmt.Printf("El área del círculo es: %.2f\n", area3)
}
