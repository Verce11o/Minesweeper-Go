package main

import (
	"minesweeper/game"
)

func main() {

	player := game.Player{Flags: game.BombAmount}
	player.NewGame()

}
