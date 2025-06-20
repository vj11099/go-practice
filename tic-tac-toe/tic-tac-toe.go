package main

import (
	game "github.com/vj11099/go-practice/tic-tac-toe/assets"
)

func main() {
	player_1 := game.MakePlayer("Sea-Salt", "x")
	player_2 := game.MakePlayer("Mark", "o")

	board := game.Reset(3)

	game.Game(player_1, player_2, board)
}
