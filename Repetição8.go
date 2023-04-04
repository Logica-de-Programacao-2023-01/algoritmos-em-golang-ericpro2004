package main

import "fmt"

func main() {
	//Divisores de um número qualquer.
	var numE int
	fmt.Println("Digite um número")
	fmt.Scanln(&numE)

	fmt.Println("Segue lista de divisores do número selecionado: ")
	for i := 1; i <= i; i++ {
		if numE%i == 0 {
			fmt.Println(i, " ")
		}
	}
}
