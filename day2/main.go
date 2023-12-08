package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	
	var re = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
//	var re = regexp.MustCompile(`(\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)
	
	var digitMap = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	
	
	
	var res [][]string
	var firstValue = 0
	var secondValue = 0
	total := 0
	
	calibrationValues, err := os.Open("../datasets/calibration_values.txt")
	if err != nil {
		log.Printf("error occurred reading file: %v", err.Error())
		return
	}
	
	scanner := bufio.NewScanner(calibrationValues)
	for scanner.Scan() {
		res = re.FindAllStringSubmatch(scanner.Text(), -1)
		if len(res) > 1 {
			if _, ok := digitMap[res[0][0]]; ok {
				firstValue = digitMap[res[len(res)-1][0]]
			} else {
				value, _ := strconv.Atoi(res[0][0])
				firstValue = value
			}
			if _, ok := digitMap[res[len(res)-1][0]]; ok {
				secondValue = digitMap[res[len(res)-1][0]]
			} else {
				value, _ := strconv.Atoi(res[len(res)-1][0])
				secondValue = value
			}
			value, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstValue, secondValue))
			total += value
		} else {
			if _, ok := digitMap[res[0][0]]; ok {
				firstValue = digitMap[res[len(res)-1][0]]
			} else {
				value, _ := strconv.Atoi(res[0][0])
				firstValue = value
			}
			value, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstValue, secondValue))
			total += value		}
	}
	
	log.Println(total)
}