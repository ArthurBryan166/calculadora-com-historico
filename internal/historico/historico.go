package historico

import (
	"fmt"
	"github.com/ArthurBryan166/Projetos/calculadora/cmd/commands"
)

type Historico struct {
	Num1      float64
	Num2      float64
	Resultado float64
	Operador  string
}

var operacoes []Historico

// adiciona o cálculo em um slice do tipo struct (Histórico)
func HistoricoCalc(n1 float64, n2 float64, res float64, opr string) {
	operacoes = append(operacoes, Historico{Num1: n1, Num2: n2, Resultado: res, Operador: opr})
}

// mostra o histórico no terminal
func MostrarHistorico() {
	commands.Clear()
	for _, op := range operacoes {
		fmt.Printf("%v %v %v = %v\n", op.Num1, op.Operador, op.Num2, op.Resultado)
	}
	commands.Pause()
	commands.Clear()
}