package main

import "fmt"

func main() {
	//Ordem dos números de 1 a 10.
	fmt.Println("A ordem crescente dos números é: ")
	for i := 1; i < 11; i++ {
		if i == 11 {
			break
		}
		fmt.Println(i)
	}
}
