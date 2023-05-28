package main

import "fmt"

func main() {
	//Leitura de vários números e comparação de grandeza entre eles.
	var numero, maior int

	fmt.Println("Digite uma sequência de números inteiros (digite 0 para parar):")
	fmt.Print("Número: ")
	fmt.Scanln(&numero)
	maior = numero
	
	for numero != 0 {
		if numero > maior {
			maior = numero
		}
		fmt.Print("Número: ")
		fmt.Scanln(&numero)
	}
	fmt.Println("O maior número é:", maior)
}
