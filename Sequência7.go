package main

import "fmt"

func main() {
	//Antecessor e sucessor de um número.
	var numqualquer int

	fmt.Print("Digite um número qualquer: ")
	fmt.Scanln(&numqualquer)

	//antecessor e sucessor
	antecessor := numqualquer - 1
	sucessor := numqualquer + 1
	println("O antecessor do número é: ", antecessor, "O sucessor do número é: ", sucessor)
}
