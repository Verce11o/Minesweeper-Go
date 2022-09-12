package main

import (
	"github.com/Verce11o/Minesweeper-Go/game"
)

func main() {

	player := game.Player{Flags: game.BombAmount}
	player.NewGame()

}
