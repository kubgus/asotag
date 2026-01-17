package game

import "fmt"

type EntityActive interface {
	// Used to perform any setup or actions before the entity's turn begins.
	BeforeTurn(context *Context)
	// Used to define the actions taken by the entity during its turn.
	OnTurn(context *Context) (string, bool)
}

func Input(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Print(ColTooltip("> "))
	fmt.Scan(&input)
	fmt.Println()
	return input
}
