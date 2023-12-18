package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	
	sum := 0
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
	for scanner.Scan() {
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
		
		sum += intersect(sliceOne, sliceTwo)
	}
	
	log.Println(time.Now().Sub(runTime))
	log.Println(sum)
}

func intersect(sliceOne []int, sliceTwo []int) int {
	var result = 0
	var intersection = make([]int, 0)
	
	for _, sOne := range sliceOne {
		for _, sTwo := range sliceTwo {
			if sOne == sTwo {
				intersection = append(intersection, sOne)
				break
			}
		}
	}
	
	if len(intersection) > 0 {
		if len(intersection) == 1 {
			result = 1
		} else {
			result = int(math.Pow(2, float64(len(intersection)-1)))
		}
	}
	
	log.Println(intersection)
	return result
}