package main

import "fmt"

type game struct {
	state []([]uint8)
	turn  bool
	count uint8
}

func newGame() *game {
	return &game{
		state: []([]uint8){
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
		turn:  true,
		count: uint8(0),
	}
}

func (g *game) printGame() {
	for _, row := range g.state {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print("_")
			} else if cell == 1 {
				fmt.Print("X")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func (g *game) printTurn() {
	if g.turn {
		fmt.Println("It's X's turn.")
	} else {
		fmt.Println("It's O's turn.")
	}
}

func (g *game) isValidMove(row, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	if g.state[row][col] != 0 {
		return false
	}
	return true
}

func (g *game) makeMove(row, col int) bool {
	var isValid = g.isValidMove(row, col)
	if isValid {
		if g.turn {
			g.state[row][col] = 1
		} else {
			g.state[row][col] = 2
		}
		g.turn = !g.turn
		g.count++
	}
	return isValid
}

func hasWon(state []([]uint8)) bool {
	for _, row := range state {
		if row[0] == row[1] && row[1] == row[2] && row[0] != 0 {
			return true
		}
	}
	for _, col := range state {
		if col[0] == col[1] && col[1] == col[2] && col[0] != 0 {
			return true
		}
	}
	if state[0][0] == state[1][1] && state[1][1] == state[2][2] && state[0][0] != 0 {
		return true
	}
	if state[0][2] == state[1][1] && state[1][1] == state[2][0] && state[0][2] != 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")

	g := newGame()

	for true {
		g.printGame()
		g.printTurn()
		var row, col int
		fmt.Print("Enter row: ")
		fmt.Scanf("%d", &row)
		fmt.Print("Enter col: ")
		fmt.Scanf("%d", &col)
		for true {
			if g.makeMove(row, col) {
				break
			}
			fmt.Print("Invalid move. Try again.\n")
			fmt.Print("Enter row: ")
			fmt.Scanf("%d", &row)
			fmt.Print("Enter col: ")
			fmt.Scanf("%d", &col)
		}
		if hasWon(g.state) {
			var player string
			if g.turn {
				player = "O"
			} else {
				player = "X"
			}
			g.printGame()
			fmt.Printf("%s has won!\n", player)
			return
		} else if g.count == 9 {
			g.printGame()
			fmt.Println("It's a draw!")
			return
		}
	}
}
