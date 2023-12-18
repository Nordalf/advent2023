package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)


func main() {

	sum := 0
	var scratchcardsMap = make(map[string]int)
	scratchValues, err := os.Open("../../datasets/scratchcards.txt")
	if err != nil {
		log.Printf("error occurred reading file: %v", err.Error())
		return
	}

	defer func(scratchValues *os.File) {
		closeErr := scratchValues.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}(scratchValues)

	scanner := bufio.NewScanner(scratchValues)
	var splittedValues []string
	var splittedNumbers []string
	var splittedSliceOne []string
	var sliceOne = make([]int, 0)
	var splittedSliceTwo []string
	var sliceTwo = make([]int, 0)

	runTime := time.Now()
	numberOfLines := 0
	for scanner.Scan() {
		numberOfLines++
		// KEEP ALLOCATIONS BUT CLEAR ARRAY
		splittedValues = splittedValues[:0]
		splittedNumbers = splittedNumbers[:0]
		splittedSliceOne = splittedSliceOne[:0]
		sliceOne = sliceOne[:0]
		splittedSliceTwo = splittedSliceTwo[:0]
		sliceTwo = sliceTwo[:0]


		splittedValues = strings.Split(scanner.Text(), ":")

		if len(splittedValues) == 2 {
			splittedNumbers = strings.Split(splittedValues[1], "|")
			// 0 == winning numbers, 1 == numbers you have
			splittedSliceOne = strings.Fields(splittedNumbers[0])
			for _, s := range splittedSliceOne {
				temp, _ := strconv.Atoi(s)
				sliceOne = append(sliceOne, temp)
			}

			splittedSliceTwo = strings.Fields(splittedNumbers[1])
			for _, s := range splittedSliceTwo {
				temp, _ := strconv.Atoi(s)
				sliceTwo = append(sliceTwo, temp)
			}
		}

		intersections := intersect(sliceOne, sliceTwo)
		cardSplit := strings.Fields(splittedValues[0])
		cardNumber, _ := strconv.Atoi(cardSplit[1])
		
		if _, ok := scratchcardsMap[fmt.Sprintf("Card %d", cardNumber)]; ok {
			scratchcardsMap[fmt.Sprintf("Card %d", cardNumber)] = scratchcardsMap[fmt.Sprintf("Card %d", cardNumber)] + 1 // +1 for the original scratchcard
			currentInstances := scratchcardsMap[fmt.Sprintf("Card %d", cardNumber)]
			for i := 1; i <= intersections; i++ {
				if _, nextOk := scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)]; nextOk {
					scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)] = scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)] + currentInstances
				} else {
					scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)] = currentInstances
				}
			}
		} else {
			scratchcardsMap[fmt.Sprintf("Card %d", cardNumber)] = 1
			for i := 1; i <= intersections; i++ {
				if _, nextOk := scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)]; nextOk {
					scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)] = scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)]
				} else {
					scratchcardsMap[fmt.Sprintf("Card %d", cardNumber + i)] = 1
				}
			}
		}
	}
	
	for _, v := range scratchcardsMap {
		sum += v
	}
	
	log.Println(time.Now().Sub(runTime))
	log.Println(sum)
}

func intersect(sliceOne []int, sliceTwo []int) int {
	var intersection = make([]int, 0)

	for _, sOne := range sliceOne {
		for _, sTwo := range sliceTwo {
			if sOne == sTwo {
				intersection = append(intersection, sOne)
				break
			}
		}
	}
	
	return len(intersection)
}