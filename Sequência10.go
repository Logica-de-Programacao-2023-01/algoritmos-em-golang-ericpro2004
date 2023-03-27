package main

import "fmt"

func main() {
	//Conversão de quilos para libras.
	var peso float64

	fmt.Print("Digite o peso a ser convertido:")
	fmt.Scanln(&peso)

	//Cálculo de conversão.
	Cálculo := peso * 2.20462
	fmt.Printf("O valor convertido é:%.2f", Cálculo)
}
