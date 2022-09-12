package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// key
// Empty square: 0
// Snake head:   2
// Food:         5

type Game struct {
	Rows int
	Columns int
	Board [][]int
	SnakeHead [2]int
	FoodPosition [2]int
}

func main() {
	game := Game{}

	game.Rows = 15
	game.Columns = 17

	game.Board = game.setupBoard();
	game.SnakeHead[0] = 0
	game.SnakeHead[1] = 0

	game.placeFood()

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

func (game Game) placeFood() {
	possibleXY := []string{}

	for x := range game.Board {
		for y := range game.Board[x] {
			if (game.Board[x][y] == 0) {
				possibleXY = append(possibleXY, fmt.Sprintf("%d, %d", x, y))
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	lowerBoundary := 0
	higherBoundary := len(possibleXY) - 1
	randomPosition := lowerBoundary + rand.Intn(higherBoundary - lowerBoundary + 1)

	value := strings.Split(possibleXY[randomPosition], ", ")
	pX, _ := strconv.Atoi(value[0])
	pY, _ := strconv.Atoi(value[1])

	game.FoodPosition[0] = pX
	game.FoodPosition[1] = pY
	game.Board[pX][pY] = 5
}