package main

import "fmt"

func main() {

	//IMC a partir do peso e altura.
	var peso float64
	var altura float64

	fmt.Print("qual é o seu peso?:")
	fmt.Scanln(&peso)
	fmt.Print("qual é sua altura?:")
	fmt.Scanln(&altura)

	//Cálculo do IMC
	Cálculo := peso / (altura * altura)
	fmt.Printf("O seu IMC é:%.2f ", Cálculo)
}
