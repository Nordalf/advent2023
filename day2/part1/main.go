package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	sumGameIDs := 0
	
	calibrationValues, err := os.Open("../../datasets/cube_game.txt")
	if err != nil {
		log.Printf("error occurred reading file: %v", err.Error())
		return
	}
	
	defer func(calibrationValues *os.File) {
		closeErr := calibrationValues.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}(calibrationValues)
	
	scanner := bufio.NewScanner(calibrationValues)
	for scanner.Scan() {
		gameId, inputString := startGame(scanner.Text())
		sumGameIDs += playRound(gameId, inputString)
	}
	
	log.Println(sumGameIDs)
}

// startGame finds and return the gameID a long with the actual playRound
func startGame(input string) (int, string) {
	// could use regex but going with splits
	line := strings.Split(input, ":")
	game := strings.Split(line[0], " ")
	gameID, _ := strconv.Atoi(game[1])
	return gameID, line[1]
}

// playRound returns the gameId if the game possible
func playRound(gameId int, input string) int {
	var game = make(map[string]int)
	var singlePlay []string
	var cubesAndColor []string
	isPossible := true
	hand := strings.Split(input, ";")
	
	for _, val := range hand {
		singlePlay = strings.Split(val, ",")
		for _, sp := range singlePlay {
			sp = strings.TrimLeft(sp, " ")
			cubesAndColor = strings.Split(sp, " ")
			cubes, _ := strconv.Atoi(cubesAndColor[0])
			game[cubesAndColor[1]] = cubes
		}
		isPossible = isGamePossible(game)
		if !isPossible {
			return 0
		}
	}
	
	return gameId
}

func isGamePossible(game map[string]int) bool {
	if val, ok := game["red"]; ok {
		if val > 12 {
			return false
		}
	}
	
	if val, ok := game["green"]; ok {
		if val > 13 {
			return false
		}
	}
	
	if val, ok := game["blue"]; ok {
		if val > 14 {
			return false
		}
	}
	
	return true
}