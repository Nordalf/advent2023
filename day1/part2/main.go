package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	
	var re = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	
	var res [][]string
	var firstValue = 0
	var secondValue = 0
	total := 0
	
	calibrationValues, err := os.Open("../../datasets/calibration_values.txt")
	if err != nil {
		log.Printf("error occurred reading file: %v", err.Error())
		return
	}
	defer func(calibrationValues *os.File) {
		err := calibrationValues.Close()
		if err != nil {
			log.Println(err)
		}
	}(calibrationValues)
	
	var tempText string
	scanner := bufio.NewScanner(calibrationValues)
	for scanner.Scan() {
		tempText = overlaps(scanner.Text())
		res = re.FindAllStringSubmatch(tempText, -1)
		if len(res) > 1 {
			firstValue, _ = strconv.Atoi(res[0][0])
			secondValue, _ = strconv.Atoi(res[len(res)-1][0])
			value, err := strconv.Atoi(fmt.Sprintf("%d%d", firstValue, secondValue))
			if err != nil {
				log.Println(err)
			}
			total += value
		} else {
			if len(res) != 0 {
				firstValue, _ = strconv.Atoi(res[0][0])
				value, err := strconv.Atoi(fmt.Sprintf("%d%d", firstValue, firstValue))
				if err != nil {
					log.Println(err)
				}
				total += value
			}
		}
	}
	
	log.Println(total)
}

func overlaps(input string) string {
	tempText := strings.ReplaceAll(input, "one", "o1ne")
	tempText = strings.ReplaceAll(tempText, "two", "t2wo")
	tempText = strings.ReplaceAll(tempText, "three", "t3hree")
	tempText = strings.ReplaceAll(tempText, "four", "f4our")
	tempText = strings.ReplaceAll(tempText, "five", "f5ive")
	tempText = strings.ReplaceAll(tempText, "six", "s6ix")
	tempText = strings.ReplaceAll(tempText, "seven", "s7even")
	tempText = strings.ReplaceAll(tempText, "eight", "e8ight")
	tempText = strings.ReplaceAll(tempText, "nine", "n9ine")

	return tempText
}