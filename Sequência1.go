package main

import "fmt"

func main() {

	//Soma de três números inteiros.
	var num1 int
	var num2 int
	var num3 int
	fmt.Print("digite um número (inteiros):")
	fmt.Scanln(&num1)
	fmt.Print("digite outro número:")
	fmt.Scanln(&num2)
	fmt.Print("digite outro número:")
	fmt.Scanln(&num3)

	//Cálculo da soma e resposta.
	Cálculo := num3 + num2 + num1
	println("Os números são: ", num1, num2, num3, "A soma dos números é igual: ", Cálculo)
}
