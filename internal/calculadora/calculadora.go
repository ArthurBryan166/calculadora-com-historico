package calculadora

import (
	"github.com/ArthurBryan166/Projetos/calculadora/internal/historico"
	"github.com/ArthurBryan166/Projetos/calculadora/internal/operacoes"
	"github.com/ArthurBryan166/Projetos/calculadora/cmd/commands"
	"fmt"
)

func Calculadora(num1, num2 float64, operador string) {
	commands.Clear()
	var resultado float64

	switch operador {
	case "+":
		resultado = operacoes.Adicao(num1, num2)
	case "-":
		resultado = operacoes.Subtracao(num1, num2)
	case "*":
		resultado = operacoes.Multiplicacao(num1, num2)
	case "/":
		resultado = operacoes.Divisao(num1, num2)
	}

	// manda o num1, operador, num2 e resultado para o historico
	historico.HistoricoCalc(num1, num2, resultado, operador)

	fmt.Println("\nResultado:", resultado)
	commands.Pause()
	commands.Clear()
}
