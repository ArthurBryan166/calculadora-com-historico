package ui

import (
	"fmt"
	"github.com/ArthurBryan166/Projetos/calculadora/internal/historico"
	"github.com/ArthurBryan166/Projetos/calculadora/cmd/commands"
	"github.com/ArthurBryan166/Projetos/calculadora/internal/calculadora"
)

func MostrarUI(){
	var num1, num2 float64
	var operador string
	sair := 1

	// looping de uso da calculadora
	for sair != 0 {
		commands.Clear()
		var digiteAqui int
		fmt.Print("Bem vindo à calculadora\n\n1 - Calculadora\n2 - Histórico\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)
		commands.Clear()

		switch digiteAqui {
		case 0:
			sair = 0
			fmt.Println("Obrigado por usar a calculadora!!")

		case 1:
			fmt.Print("Digite um número: ")
			if _, err := fmt.Scan(&num1); err != nil{
				fmt.Printf("Erro no primeiro número: %v", err)
				commands.Pause()
				continue
			}

			fmt.Print("Digite o operador: ")
			if _, err := fmt.Scan(&operador); err != nil{
				fmt.Printf("Erro no operador: %v", err)
				commands.Pause()
				continue
			}

			fmt.Print("Digite outro número: ")
			if _, err := fmt.Scan(&num2); err != nil{
				fmt.Printf("Erro no segundo número: %v", err)
				commands.Pause()
				continue
			}
			calculadora.Calculadora(num1, num2, operador)

		case 2:
			fmt.Print("Histórico:\n\n")
			historico.MostrarHistorico()
		}
	}
}