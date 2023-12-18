package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)
var symbolRe = regexp.MustCompile(`\d+`)
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
	
	
	count := 0
	runTime := time.Now()
	for row, line := range coordinates {
		for column, data := range line {
			if data != "*" {
				continue
			}
			sum +=checkAdjacents(coordinates, row, column)
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

func checkAdjacents(coords [][]string, row, column int) int {
	if row == 0 && column == 0 {
		return checkAdjacentsRec(coords, row, column, 0, 0,  1, false, 0)
	} else if column == 0 {
		return checkAdjacentsRec(coords, row, column, -1, 0,  1, false,0 )
	} else if row == 0 {
		return checkAdjacentsRec(coords, row, column, 0, -1,  1, false, 0)
	}  else {
		return checkAdjacentsRec(coords, row, column, -1, -1,  1, false, 0)
	}
}

func checkAdjacentsRec(coords [][]string, row int, column int, y int, x int, value int, skip bool, partNumbers int) int {
	currentData := ""
	
	// Intet blev fundet
	if y == 1 && x > 1 {
		return 0
	}
	
	if y > 1 {
		if partNumbers == 2 {
			return value		
		}
		return 0
	}
	
	if row + y >= len(coords) {
		if partNumbers == 2 {
			return value		
		}
		return 0
	}
	if column + x >= len(coords[row]){
		return 0
	}
	currentData = coords[row + y][column + x]
	
	if isNumber(currentData) {
		temp, _ := strconv.Atoi(currentData)

		if (value == 1 || temp != value ) && partNumbers != 2{
			value *= temp
			partNumbers++
		}
	}
	
	if x == 1 {
		y++
		if column == 0 {
			x = -1
		} else {
			x = -2
		}
		skip = false
	}
	
	
	return checkAdjacentsRec(coords, row, column, y, x+1, value, skip, partNumbers)
}

func isNumber(data string) bool {
	return symbolRe.MatchString(data)
}

//
//95496
//564876
//242374
//532644
//510860
//359856
//107166
//123255

// 2539127