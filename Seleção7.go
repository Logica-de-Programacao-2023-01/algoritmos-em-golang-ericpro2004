package main

import "fmt"

func main() {

//Aumento de salário dos funcionários em 10% ou 5%.

	var s float64
	var ns float64
	fmt.Println("Digite o salário do funcionário:")
	fmt.Scanln(&s)

	if s < 1000 {
		ns = s + s*0.1
		fmt.Println("Novo salário (aumento de 10%):", ns)
	} else {
		ns = s + s*0.05
		fmt.Println("Novo salário (aumento de 5%):", ns)
	}
}
