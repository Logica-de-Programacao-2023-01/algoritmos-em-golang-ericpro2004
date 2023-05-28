package main

import "fmt"

func main() {
	
//Dias da semana baseados em números. 
	
	var num int
	var diadaSemana string
	fmt.Println("Digite um número (inteiro) entre 1 e 7: ")
	fmt.Scanln(&num)
	
	if num == 1 {
		diadaSemana = "domingo"
	} else if num == 2 {
		diadaSemana = "segunda-feira"
	}else if num == 3 {
		diadaSemana = "terça-feira"
	}else if num == 4 {
		diadaSemana = "quarta-feira"
	} else if num == 5 {
		diadaSemana = "quinta-feira"
	} else if num == 6 {
		diadaSemana = "sexta-feira"
	} else if num == 7 {
		diadaSemana = "sábado"
	} else {
		diadaSemana = "Inválido"
	}
	fmt.Print("Dia da semana correspondente:", diadaSemana)
}
