package main

import "fmt"

func main() {
	//Média aritmética de vários números.
	var num, som, quant int
	fmt.Println("Digite os números inteiros desejados (digite 0 para encerrar): ")
	for {
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		som += num
		quant++
	}
	if quant == 0 {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		m := float64(som) / float64(quant)
		fmt.Println("A média aritmética dos números é: ", m)
	}
}
