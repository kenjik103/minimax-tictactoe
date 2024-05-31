package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func generateBoard(boardArray string) {
	var sb strings.Builder
	for i := 0; i < 9; i++ {
		sb.WriteString(string(boardArray[i]))
		if i == 2 || i == 5 || i == 8 {
			sb.WriteString("\n")
		} else {
			sb.WriteString("|")
		}
	}
	fmt.Println(sb.String())
}

// returns bool (is there a winner or not) and string (who won: player or bot)
func checkWin(boardArray string) (bool, string) {
	for i := 0; i < 9; i += 3 {
		var row string = boardArray[i : i+3]
		if row == "ooo" {
			return true, "player"
		}
		if row == "xxx" {
			return true, "bot"
		}
	}
	for i := 0; i < 3; i++ {
		var col string = string(boardArray[i]) + string(boardArray[i+3]) + string(boardArray[i+6])
		if col == "ooo" {
			return true, "player"
		}
		if col == "xxx" {
			return true, "bot"
		}
	}
	var cross1 string = string(boardArray[0]) + string(boardArray[4]) + string(boardArray[8])
	var cross2 string = string(boardArray[2]) + string(boardArray[4]) + string(boardArray[6])

	if cross1 == "ooo" || cross2 == "ooo" {
		return true, "player"
	}
	if cross1 == "xxx" || cross2 == "xxx" {
		return true, "bot"
	}
	return false, ""

}

// extention of checkWin func; returns true if board is full or there is a winner
func checkGameOver(boardArray string) bool {
	var isWinner, _ = checkWin(boardArray)
	if isWinner {
		return true
	}
	for _, c := range boardArray {
		if c != 'x' && c != 'o' {
			return false
		}
	}
	return true
}

/*
* used for static evaluation in minimax algorithm.
Returns number empty spaces in a board. Useful because the more empty
squares, the better/worse of a loss/win it was. Allows AI make moves that win
ASAP, or lose in the most amount of moves possible if passed a rigged board.
*
*/
func getUtility(boardArray string) int {
	utility := 1
	for _, c := range boardArray {
		if c != 'x' && c != 'o' {
			utility++
		}
	}
	return utility
}

func minimax(boardArray string, depth int, maximizingPlayer bool) int {
	if checkGameOver(boardArray) {
		var win, winner = checkWin(boardArray)
		if win && winner == "bot" {
			return getUtility(boardArray)
		}
		if win && winner == "player" {
			return -(getUtility(boardArray))
		}
		return 0
	}
	if maximizingPlayer {
		var maxEval float64 = math.Inf(-1)
		var children strings.Builder
		for _, c := range boardArray {
			if c != 'x' && c != 'o' {
				children.WriteString(string(c))
			}
		}
		for _, position := range children.String() {
			updatedBoard := strings.ReplaceAll(boardArray, string(position), "x")
			var eval int = minimax(updatedBoard, depth+1, false)
			maxEval = max(maxEval, float64(eval))
		}
		return int(maxEval)
	}
	if !maximizingPlayer {
		var minEval float64 = math.Inf(1)
		var children strings.Builder
		for _, c := range boardArray {
			if c != 'x' && c != 'o' {
				children.WriteString(string(c))
			}
		}
		for _, position := range children.String() {
			updatedBoard := strings.ReplaceAll(boardArray, string(position), "o")
			var eval int = minimax(updatedBoard, depth+1, true)
			minEval = min(minEval, float64(eval))
		}
		return int(minEval)
	}
	return 0
}

// calculates best move using minimax algorithm, passes that position out as a string
func botMove(boardArray string) string {
	var positions strings.Builder
	for _, c := range boardArray {
		if c != 'x' && c != 'o' {
			positions.WriteString(string(c))
		}
	}
	bestEval := math.Inf(-1)
	var bestMove string
	for _, move := range positions.String() {
		modifiedBoard := strings.ReplaceAll(boardArray, string(move), "x")
		var eval int = minimax(modifiedBoard, 0, false)
		if float64(eval) > bestEval {
			bestEval = float64(eval)
			bestMove = string(move)
		}
	}
	return bestMove
}

func gameplayLoop() {
	boardArray := "123456789"
	var position string

	for {
		//take and process user input
		generateBoard(boardArray)
		fmt.Println("Enter a position (1-9)")
		fmt.Scanln(&position)

		i, e := strconv.Atoi(position)
		if e != nil || i < 0 || i > 9 || !strings.Contains(boardArray, position) {
			fmt.Println("Invalid Input")
			continue
		}

		//print resulting board
		fmt.Println("Your Move:")
		boardArray = strings.ReplaceAll(boardArray, string(position), "o")

		generateBoard(boardArray)

		//check if user move resulted in game over
		if checkGameOver(boardArray) {
			fmt.Println("Game Over!")
			break
		}

		//calculate and display bot's move
		fmt.Println("Bot's move: ")

		boardArray = strings.ReplaceAll(boardArray, botMove(boardArray), "x")

		//check if bot's move resulted in game over
		if checkGameOver(boardArray) {
			generateBoard(boardArray)
			fmt.Println("Game Over!")
			break
		}

	}
}

func main() {
	gameplayLoop()
}
