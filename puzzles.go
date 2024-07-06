package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

// Define the structure for the JSON data
type Puzzle struct {
    FEN             string  `json:"FEN"`
    FirstToMove     string  `json:"first_to_move"`
    GameURL         string  `json:"GameUrl"`
    Moves           string  `json:"Moves"`
    NbPlays         int     `json:"NbPlays"`
    OpeningTags     *string `json:"OpeningTags"`
    Popularity      int     `json:"Popularity"`
    Rating          int     `json:"Rating"`
    RatingBy100     int     `json:"Rating_by100"`
    RatingDeviation int     `json:"RatingDeviation"`
    Themes          string  `json:"Themes"`
}

func getPuzzlesByRating(elo string) map[string]Puzzle {
	fileName := fmt.Sprintf("puzzles/%s.json", elo)
    // Open the JSON file
    jsonFile, err := os.Open(fileName)
    if err != nil {
        log.Fatalf("Failed to open JSON file: %s", err)
    }
    defer jsonFile.Close()

    // Read the JSON file
    byteValue, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        log.Fatalf("Failed to read JSON file: %s", err)
    }

    // Create a map to hold the unmarshaled data
    var puzzles map[string]Puzzle

    // Unmarshal the JSON data into the map
    err = json.Unmarshal(byteValue, &puzzles)
    if err != nil {
        log.Fatalf("Failed to unmarshal JSON: %s", err)
    }

	return puzzles
}