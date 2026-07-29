package commands

import "fmt"

func Clear(){
	// limpar terminal
	fmt.Print("\033[H\033[2J")
}