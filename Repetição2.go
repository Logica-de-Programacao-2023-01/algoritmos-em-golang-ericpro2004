package main

import "fmt"

func main() {
	//Números pares de 0 a 20.
	fmt.Println("Os números são: ")
	for i := 0; i < 21; i++ {
		if i == 1 {
			continue
		}
		if i == 3 {
			continue
		}
		if i == 5 {
			continue
		}
		if i == 7 {
			continue
		}
		if i == 9 {
			continue
		}
		if i == 11 {
			continue
		}
		if i == 13 {
			continue
		}
		if i == 15 {
			continue
		}
		if i == 17 {
			continue
		}
		if i == 19 {
			continue
		}

		fmt.Println(i)
	}
}
