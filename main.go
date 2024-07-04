package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {
	fenString := getFENString()
	cb := NewFENChessboard(fenString)
	fmt.Println("\n" + cb.GetFormattedBoard())
}
