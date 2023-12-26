package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Boatrace struct {
	Time int
	Duration int
}

func main() {

	var boatrace Boatrace

	races, err := os.Open("../../datasets/boatraces.txt")
	if err != nil {
		log.Printf("error occurred reading file: %v", err.Error())
		return
	}

	defer func(races *os.File) {
		closeErr := races.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}(races)

	scanner := bufio.NewScanner(races)
	
	// Time
	scanner.Scan()
	timeLines := strings.Split(scanner.Text(), ":")
	// Duration
	scanner.Scan()
	durations := strings.Split(scanner.Text(), ":")
	
	
	convertedTime, _ := strconv.Atoi(strings.ReplaceAll(timeLines[1], " ", ""))
	convertedDuration, _ := strconv.Atoi(strings.ReplaceAll(durations[1], " ", ""))
	boatrace = Boatrace {
		Time:     convertedTime,
		Duration: convertedDuration,
	}

	var result = holdButton(boatrace)

	log.Println(timeLines)
	log.Println(durations)
	log.Println(result)
}

// holdButton returns different ways to win the boatrace based on time and duration
func holdButton(boatrace Boatrace) int {
	var result = 0
	var temp = 0
	for i := 1; i < boatrace.Time; i++ {
		temp = (boatrace.Time - i) * i
		if temp > boatrace.Duration {
			result++
		}
	}

	return result
}