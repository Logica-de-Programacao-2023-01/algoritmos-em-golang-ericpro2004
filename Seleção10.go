package main

import "fmt"

func main() {
//Leitura de idade.
	var idade int
	var class string
	fmt.Println("Digite a idade da pessoa:")
	fmt.Scanln(&idade)

	switch {
	case idade <= 9:
		class = "mirim"
	case idade <= 13:
		class = "infantil"
	case idade <= 17:
		class = "juvenil"
	default:
		class = "adulto"
	}
	fmt.Println("A classificação da pessoa é:", class)
}
