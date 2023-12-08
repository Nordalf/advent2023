package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
//	var calibrationValues = []string{
//		"1abc2",
//		"pqr3stu8vwx",
//		"a1b2c3d4e5f",
//		"treb7uchet",
//	}
	
	var re = regexp.MustCompile(`(\d)`)
	var res [][]string
	var consolidatedString string
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
			consolidatedString = res[0][0] + res[len(res)-1][0]
			value, _ := strconv.Atoi(consolidatedString)
			total += value
		} else {
			consolidatedString = res[0][0] + res[0][0]
			value, _ := strconv.Atoi(consolidatedString)
			total += value
		}
	}
	
	log.Println(total)
}