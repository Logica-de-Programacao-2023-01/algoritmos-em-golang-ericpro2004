package main

import "fmt"

func main() {

	//Comparação de grandeza entre 3 números.

	var num1, num2, num3 int

	fmt.Printf("Digite um número (primeiro): ")
	fmt.Scanln(&num1)
	fmt.Printf("Digite um número (segundo): ")
	fmt.Scanln(&num2)
	fmt.Printf("Digite um número (último): ")
	fmt.Scanln(&num3)

	if num1 < num2 && num1 < num3 {
		fmt.Printf("%d é o menor número entre os três números: ", num1)
	} else if num2 < num1 && num2 < num1 {
		fmt.Printf("%d é o menor número entre os três números: ", num2)
	} else {
		fmt.Printf("%d é o menor número entre os três números: ", num3)
	}
}
