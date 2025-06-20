package game

import "fmt"

func Game(player_1, player_2 Player, board Board) {
	turn := 0
	gameDone := false

	for !gameDone {
		var player Player
		if turn%2 == 0 {
			player = player_1
		} else {
			player = player_2
		}
		var x, y int32
		fmt.Print(player.name, " enter [X] position of the marker: ")
		fmt.Scan(&x)

		fmt.Print(player.name, " enter [Y] position of the marker: ")
		fmt.Scan(&y)

		gameDone = player.Place(board, x, y)
		if gameDone {
			fmt.Println("Player ", player.name, " wins")
			break
		}
		turn++
	}
}
