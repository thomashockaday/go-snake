package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	Rows int
	Columns int
	Board [][]int
	SnakeHead [2]int
	FoodPosition [2]int
	LengthOfSnake int
	Trail [][]int
}

func main() {
	game := Game{}

	game.Rows = 15
	game.Columns = 17

	game.Board = game.setupBoard()
	game.SnakeHead[0] = 0
	game.SnakeHead[1] = 0

	game.placeFood()
	game.printBoard()

	for {
		var direction string
		fmt.Print("Move (w, a, s, d): ")
		fmt.Scanln(&direction)

		err := game.move(direction)

		if (err != nil) {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		if (game.SnakeHead[0] == game.FoodPosition[0] && game.SnakeHead[1] == game.FoodPosition[1]) {
			fmt.Println("Success")

			game.placeFood()
		}

		game.printBoard()
	}
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

func (game *Game) placeFood() {
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

func (game *Game) move (direction string) error {
	if direction != "w" && direction != "a" && direction != "s" && direction != "d" {
		return nil
	}

	prevX := game.SnakeHead[0]
	prevY := game.SnakeHead[1]

	game.Board[prevX][prevY] = 0

	if (direction == "w") {
		if (game.SnakeHead[0] - 1 > -1) {
			oldX := game.SnakeHead[0]
			game.SnakeHead[0] = oldX - 1
		}
	}

	if direction == "d"{
		if game.SnakeHead[1] + 1 < game.Columns{
			oldY := game.SnakeHead[1]
			game.SnakeHead[1] = oldY + 1
		}
	}

	if (direction == "s") {
		if (game.SnakeHead[0] + 1 < game.Rows) {
			oldX := game.SnakeHead[0]
			game.SnakeHead[0] = oldX + 1
		}
	}

	if direction == "a"{
		if game.SnakeHead[1] - 1 > -1{
			oldY := game.SnakeHead[1]
			game.SnakeHead[1] = oldY - 1
		}
	}

	x := game.SnakeHead[0]
	y := game.SnakeHead[1]

	if (game.Board[x][y] == 1) {
		return errors.New("Game over")
	}

	game.Board[x][y] = 2
	game.Board[prevX][prevY] = 1
	game.Trail = append(game.Trail, []int{prevX, prevY})
	return nil
}

func (game *Game) cleanTrail() {
	game.LengthOfSnake += len(game.Trail)

	for coordinate := range game.Trail {
		xy := game.Trail[coordinate]
		game.Board[xy[0]][xy[1]] = 1
	}

	game.Trail = [][]int{}
}