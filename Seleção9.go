package main

import (
	"fmt"
	"sort"
)

func main() {
	
//Três números em ordem crescente.
	
	var num1, num2, num3 float64

	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Println("Digite o terceiro número: ")
	fmt.Scanln(&num3)

	nums := []float64{num1, num2, num3}
	sort.Float64s(nums)

	fmt.Println("Ordem crescente dos números: ", nums)
}
