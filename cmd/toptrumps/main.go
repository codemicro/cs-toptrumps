package main

import (
	"fmt"
	"time"

	"github.com/codemicro/cs-toptrumps/internal/cards"
	"github.com/codemicro/cs-toptrumps/internal/game"
	"github.com/codemicro/cs-toptrumps/internal/helpers"
)

func main() {

	helpers.ClearConsole()

	fmt.Printf("Top trumps, but it's planes and only has %d cards\n", len(cards.AllCards))
	fmt.Println()
	time.Sleep(time.Second)

	g := game.New(cards.SplitCards(game.NumPlayers))
	g.Run()
}
