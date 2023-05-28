package main

import "fmt"

func main() {

	// Multiplicação e soma baseadas na positividade ou negatividade de dois números.

	var num1 int
	var num2 int
	var tipocal int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	if num1 > 0 && num2 > 0 {
		tipocal = num1 * num2
		fmt.Print("Resultado da multiplicação: ", tipocal)
	} else {
		tipocal = num1 + num2
		fmt.Print("Resultado da soma: ", tipocal)
	}
}
