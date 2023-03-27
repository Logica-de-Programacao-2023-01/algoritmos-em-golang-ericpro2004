package main

import (
	"fmt"
)

func main() {

	//Comparação de grandeza entre números

	var num1, num2 int

	fmt.Print("Digite o número a ser comparado: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número a ser comparado: ")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Printf("%d é maior que %d: ", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%d é menor que %d: ", num1, num2)
	} else {
		fmt.Printf("%d é igual a %d", num1, num2)
	}
}
