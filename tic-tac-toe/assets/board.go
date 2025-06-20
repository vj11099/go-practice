package game

import "fmt"

type Board struct {
	size         int32
	rowCounts    map[int32]map[string]int
	columnCounts map[int32]map[string]int
	diagCounts   map[string]map[string]int
	board        [][]string
}

func Reset(size int32) Board {
	boardState := make([][]string, size)
	for i := range boardState {
		boardState[i] = make([]string, size)
	}
	return Board{
		size,
		make(map[int32]map[string]int),
		make(map[int32]map[string]int),
		make(map[string]map[string]int),
		boardState,
	}
}

func (player Player) Place(board Board, x, y int32) bool {
	if x >= board.size || y >= board.size {
		panic("Input size exceeded the board size")
	}
	if board.board[x][y] != "" {
		panic("Illegal move")
	}

	board.board[x][y] = player.marker

	// row counts
	if board.rowCounts[y] == nil {
		board.rowCounts[y] = make(map[string]int)
		if board.rowCounts[y][player.marker] == 0 {
			board.rowCounts[y][player.marker] = board.rowCounts[y][player.marker] + 1
		}
	} else {
		board.rowCounts[y][player.marker] = board.rowCounts[y][player.marker] + 1
	}

	// column counts
	if board.columnCounts[x] == nil {
		board.columnCounts[x] = make(map[string]int)
		if board.columnCounts[x][player.marker] == 0 {
			board.columnCounts[x][player.marker] = board.columnCounts[x][player.marker] + 1
		}
	} else {
		board.columnCounts[x][player.marker] = board.columnCounts[x][player.marker] + 1
	}

	// diagonal counts
	if x == y {
		if board.diagCounts["fwd"] == nil {
			board.diagCounts["fwd"] = make(map[string]int)
			if board.diagCounts["fwd"][player.marker] == 0 {
				board.diagCounts["fwd"][player.marker] = board.diagCounts["fwd"][player.marker] + 1
			}
		} else {
			board.diagCounts["fwd"][player.marker] = board.diagCounts["fwd"][player.marker] + 1
		}
	}
	if x+y+1 == board.size {
		if board.diagCounts["bwd"] == nil {
			board.diagCounts["bwd"] = make(map[string]int)
			if board.diagCounts["bwd"][player.marker] == 0 {
				board.diagCounts["bwd"][player.marker] = board.diagCounts["bwd"][player.marker] + 1
			}
		} else {
			board.diagCounts["bwd"][player.marker] = board.diagCounts["bwd"][player.marker] + 1
		}
	}

	// print board
	for i := range board.board {
		for j := range board.board[i] {
			if j == 0 {
				fmt.Print("|")
			}
			if board.board[i][j] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(board.board[i][j])
			}
			if j == int(board.size)-1 {
				fmt.Println("|")
			}
		}
	}

	fmt.Println("diagCounts", board.diagCounts["bwd"])
	if board.columnCounts[x][player.marker] == int(board.size) ||
		board.rowCounts[y][player.marker] == int(board.size) ||
		board.diagCounts["fwd"][player.marker] == int(board.size) ||
		board.diagCounts["bwd"][player.marker] == int(board.size) {
		return true
	}
	return false
}
