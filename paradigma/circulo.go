package main

import (
	"fmt"
	"math"
	"strconv"
)

func circulo() {
	var radio float64 = 0
	fmt.Print("Ingrese el radio:")
	fmt.Scanf("%f\n", &radio)
	if radio < 0 {
		fmt.Print("El numero es menor a 10")
	} else {
		area := math.Pi * math.Pow(radio, 2)
		fmt.Printf("El área del círculo es: %.2f\n", area)
		if area < 10 {
			fmt.Print("El area es menor a 10")
		} else if area >= 10 {
			fmt.Print("El Area es mayor a 10")
		} else {
			fmt.Print("El Area es igual idk")
		}
	}
	var valor1 = "53"
	var valor2 = 255
	valor3, err := strconv.Atoi(valor1)
	if err != nil {
		fmt.Print("\n Hay un error, no se pudo convertir ")
	} else {
		fmt.Print("\n El valor convertido es ", valor3)
	}
	valor4, err := strconv.Itoa(valor2)
	if err != nil {
		fmt.Print("\n Hay un error, no se pudo convertir ")
	} else {
		fmt.Print("\n El valor convertido es ", valor4)
	}
}

func main() {
	circulo()
}
