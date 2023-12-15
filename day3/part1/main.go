package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)
var symbolRe = regexp.MustCompile(`[^.\d\n;]+`)
var coordRe = regexp.MustCompile(`\d+|.`)
var lenghtOfLongestLine = 0

func main() {
	
	sum := 0
	var coordinates = make([][]string, 0)
	
	calibrationValues, err := os.Open("../../datasets/engine_schematic.txt")
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
		setupCoordinates(scanner.Text(), &coordinates)
	}
	
	// Set all line lengths to be equal by appending dots
	for lineIndex, line := range coordinates {
		for i := 0; i < lenghtOfLongestLine - len(line); i++ {
			coordinates[lineIndex] = append(coordinates[lineIndex], ".")
		}
	}
	
	
	shouldAppend := false
	skip := false
	count := 0
	runTime := time.Now()
	for row, line := range coordinates {
		skip = false
		for column, data := range line {
			shouldAppend = false
			val, convErr := strconv.Atoi(data)
			if convErr != nil {
				skip = false
				continue
			}
			
			// Must be a number
			if skip {
				continue
			}
			
//			log.Printf("VALUE FOUND: %d ROW: %d, COLUMN: %d ", val, row, column)
			if row == 0 && column == 0 {
				shouldAppend = checkAdjacents(coordinates, row, column, 0, 0, false)
			} else if column == 0 {
				shouldAppend = checkAdjacents(coordinates, row, column, -1, 0, false)
			} else if row == 0 {
				shouldAppend = checkAdjacents(coordinates, row, column, 0, -1, false)
			}  else {
				shouldAppend = checkAdjacents(coordinates, row, column, -1, -1, false)
			}
			
			if shouldAppend {
				count++
				sum += val
				skip = true
			}
		}
	}

	log.Println(time.Now().Sub(runTime))
	log.Println(sum)
	log.Println(count)
}

func setupCoordinates(row string, coordinates *[][]string) {
	var found = coordRe.FindAllString(row, -1)
	
	var temp []string
	var convErr error
	loopAmount := 0
	for _, f := range found {
		loopAmount = len(f) - 1
		_, convErr = strconv.Atoi(f)
		if convErr != nil {
			temp = append(temp, f)
		} else {
			temp = append(temp, f)
			for la := 0; la < loopAmount; la++ {
				temp = append(temp, f)
			}
		}
	}

	*coordinates = append(*coordinates, temp)
}

func checkAdjacents(coords [][]string, row, column, y, x int, found bool) bool {
	
	if found {
		return true
	}
	currentData := ""
	
	// Intet blev fundet
	if y == 1 && x > 1 {
		return false
	}
	
	if y > 1 {
		return false
	}
	
	if row + y >= len(coords) {
		return false
	}
	if column + x >= len(coords[row]){
		return false
	}
	currentData = coords[row + y][column + x]
	if x == 1 {
		y++
		if column == 0 {
			x = -1
		} else {
			x = -2
		}
	}
	
	return checkAdjacents(coords, row, column, y, x+1, isSymbol(currentData))
}

func isSymbol(data string) bool {
	return symbolRe.MatchString(data)
}
