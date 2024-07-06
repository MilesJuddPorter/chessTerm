package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"flag"
	"strconv"
)

func validateFENString(fenString string) bool {
	// Check if the FEN string has 8 rows - slash separated
	if len(strings.Split(fenString, "/")) != 8 {
		fmt.Println("Invalid FEN string. Please enter a valid FEN string.")
		return false
	} else {
		return true
	}
}

func getFENString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter FEN string: ")
	fenString, _ := reader.ReadString('\n')

	if !validateFENString(fenString) {
		return getFENString()
	} else {
		return fenString
	}
}


// // Main for the FEN Chessboard input
// func main() {
// 	fenString := getFENString()
// 	cb := NewFENChessboard(fenString)
// 	fmt.Println("\n" + cb.GetFormattedBoard())
// }
func getEloRating() string {
	// Creates an 'elo' flag that the user will input
	elo := flag.String("elo", "100", "Elo rating for the puzzles")
	flag.Parse()
	elo_numeric, _ := strconv.Atoi(*elo)

	if elo_numeric % 100 == 0 && elo_numeric >= 300 && elo_numeric <= 3000 {
		return *elo
	} else {
		fmt.Println("Invalid puzzle rating. Please enter a valid rating. [300-3000]. \ne.g. chessTerm -elo 1200")
		return "0"
	}
}

func getMove() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ['g' to give up]\n")
	move, _ := reader.ReadString('\n')
	move = strings.TrimSpace(move)
	if move != "g" && len(move) != 4 {
		fmt.Println("Invalid move. Follow the format [start square][end square] e.g. e2e4\nOr enter 'g' to give up.")
		return getMove()
	}
	return move
}

// Main for puzzles
func main() {
	elo := getEloRating()
	if elo == "0" {
		return
	}
	puzzles := getPuzzlesByRating(elo)

	var cb *Chessboard

	for _, puzzle := range puzzles {
		cb = NewFENChessboard(puzzle.FEN)	

		fmt.Println("\n" + cb.GetFormattedBoard())
		
		// The puzzle is set up before the computer moves, so this is flipped for the person
		if puzzle.FirstToMove == "w" {
			fmt.Printf("White plays %s\n", strings.Split(puzzle.Moves, " ")[0])
			fmt.Println("Black to move")
		} else {
			fmt.Printf("Black plays %s\n", strings.Split(puzzle.Moves, " ")[0])
			fmt.Println("White to move")
		}
		move := getMove()

		if move == "g" {
			fmt.Printf("Giving up. The correct moves were: %s \033[0m\n", strings.Split(puzzle.Moves, " ")[1:])
		} else {
			if move == strings.Split(puzzle.Moves, " ")[1] {
				fmt.Println("\033[32m\u2705 Correct! Well done.\033[0m")
			} else {
				fmt.Printf("\033[31m\u274C Incorrect. The correct move was %s \033[0m\n", strings.Split(puzzle.Moves, " ")[1])
			}
		}
		time.Sleep(2 * time.Second)
	}
}
