package main

import "fmt"

// key
// Empty square: 0
// Snake head:   2

type Game struct {
	Rows int
	Columns int
	Board [][]int
	SnakeHead [2]int
}

func main() {
	game := Game{}

	game.Rows = 15
	game.Columns = 17

	game.Board = game.setupBoard();
	game.SnakeHead[0] = 0
	game.SnakeHead[1] = 0

	game.printBoard()
}

func (game Game) setupBoard() [][]int {
	board := make([][]int, game.Rows)

	for i := range board {
		board[i] = make([]int, game.Columns)
	}

	board[0][0] = 2

	return board
}

func (game Game) printBoard() {
	for x := range game.Board {
		for y := range game.Board[x] {
			fmt.Print(game.Board[x][y])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Printf("Coordinate of snake head: {%d, %d}\n", game.SnakeHead[0], game.SnakeHead[1])
}