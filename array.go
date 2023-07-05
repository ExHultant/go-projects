package main

import "fmt"

func generarPiramide(nivel int) [][]int {
	piramide := make([][]int, nivel)

	for i := 0; i < nivel; i++ {
		piramide[i] = make([]int, i+1)
		piramide[i][0] = 1
		piramide[i][i] = 1

		for j := 1; j < i; j++ {
			piramide[i][j] = piramide[i-1][j-1] + piramide[i-1][j]
		}
	}

	return piramide
}

func imprimirPiramide(piramide [][]int) {
	nivel := len(piramide)

	for i := 0; i < nivel; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(piramide[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	nivel := 5
	piramide := generarPiramide(nivel)
	imprimirPiramide(piramide)
}
