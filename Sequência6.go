package main

import "fmt"

func main() {

	var salário float64

	fmt.Print("Digite seu salário atual:")
	fmt.Scanln(&salário)

	//Aumento do Salário.
	Cálculo := (salário/100)*15 + salário
	fmt.Printf("Seu salário foi reajustado em:%.2f ", Cálculo)
}
