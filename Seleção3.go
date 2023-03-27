package main

import "fmt"

func main() {

	//Verificação de números.

	var num int

	fmt.Printf("digite um número inteiro: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("%d é par", num)
	} else {
		fmt.Printf("%d é ímpar", num)
	}
}
