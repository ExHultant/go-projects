package main

import "fmt"

func slice() {
	names := []string{"Pedro", "Juan", "Carlos"}
	for i, valor := range names {
		fmt.Println("Indice", i)
		fmt.Println("Contenido", valor)
	}
	//Construir una funcion que determine que una palabra es palindromo
}
func polindromo() {
	var word string
	fmt.Print("Ingrese una palabra: ")
	fmt.Scan(&word)
	esPolindromo := true
	n := len(word)

	for i := 0; i < n/2; i++ {
		if word[i] != word[n-i-1] {
			esPolindromo = false
			break
		}
	}
	if esPolindromo {
		fmt.Print("%s es un polindromo\n", word)
	} else {
		fmt.Print("%s no es un polindromo\n", word)
	}
}

func main() {
	slice()
	polindromo()
}
