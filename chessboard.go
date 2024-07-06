package main

import (
	"strconv"
	"strings"
)

// Creates a chessboard class

type Row struct {
	Row [8]string
}

type Chessboard struct {
	Board [8]Row
}

func StartingChessboard() *Chessboard {
	cb := &Chessboard{
		Board: [8]Row{
			{Row: [8]string{"r", "n", "b", "q", "k", "b", "n", "r"}},
			{Row: [8]string{"p", "p", "p", "p", "p", "p", "p", "p"}},
			{Row: [8]string{" ", " ", " ", " ", " ", " ", " ", " "}},
			{Row: [8]string{" ", " ", " ", " ", " ", " ", " ", " "}},
			{Row: [8]string{" ", " ", " ", " ", " ", " ", " ", " "}},
			{Row: [8]string{" ", " ", " ", " ", " ", " ", " ", " "}},
			{Row: [8]string{"P", "P", "P", "P", "P", "P", "P", "P"}},
			{Row: [8]string{"R", "N", "B", "Q", "K", "B", "N", "R"}},
		},
	}
	return cb
}

// ParseFENRow parses a single FEN row string into a Row struct
func ParseFENRow(fenRow string) Row {
	if strings.Contains(fenRow, " ") {
		fenRow = strings.Split(fenRow, " ")[0]
	}
	rowVals := [8]string{}
	idxToUpdate := 0
	for _, char := range fenRow {
		if strings.ContainsAny(string(char), "12345678") {
			spacesToAdd, _ := strconv.Atoi(string(char))
			for i := 0; i < spacesToAdd; i++ {
				rowVals[idxToUpdate] = " "
				idxToUpdate += 1
			}
		} else {
			rowVals[idxToUpdate] = string(char)
			idxToUpdate += 1 
		}
	}
	return Row{Row: rowVals}
}

// NewFENChessboard initializes a Chessboard from a FEN string
func NewFENChessboard(fenString string) *Chessboard {
	fenString = strings.TrimSpace(fenString)
	fenRows := strings.Split(fenString, "/")
	boardRows := [8]Row{}
	for ii, row := range fenRows {
		boardRows[ii] = ParseFENRow(row)
	}
	cb := &Chessboard{
		Board: boardRows,
	}
	return cb
}

// GetFormattedBoard returns a formatted string representation of the chessboard
func (cb *Chessboard) GetFormattedBoard() string {
	var formattedStr string
	
	for rank, row := range cb.Board {
		for fileNum, square := range row.Row {
			if fileNum == 7 {
				formattedStr += square + "\n"
				if rank != 7 {
					formattedStr += "-----------------------------\n"
				}
			} else {
				formattedStr += square + " | "
			}
		}
	}
	reset := "\033[0m"
	white := "\033[33m"
	black := "\033[35m"
	SymbolMapper := map[string]string{
		"R": white+"\u265C"+reset,
		"N": white+"\u265E"+reset,
		"B": white+"\u265D"+reset,
		"Q": white+"\u265B"+reset,
		"K": white+"\u265A"+reset,
		"P": white+"\u265F"+reset,
		"r": black+"\u265C"+reset,
		"n": black+"\u265E"+reset,
		"b": black+"\u265D"+reset,
		"q": black+"\u265B"+reset,
		"k": black+"\u265A"+reset,
		"p": black+"\u265F"+reset,
	}

	for textChar, symbolChar := range SymbolMapper {
		formattedStr = strings.ReplaceAll(formattedStr, textChar, symbolChar)
	}

	return formattedStr
}
