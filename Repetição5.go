package main

import "fmt"

func main() {
	//Ordem dos decrescente dos números de 10 a 1.
	fmt.Println("A ordem decrescente dos números é: ")
	for i := 10; i < 11; i-- {
		if i == 0 {
			break
		}
		fmt.Println(i)
	}
}
