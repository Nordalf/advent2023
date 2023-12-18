package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Boatrace struct {
	Time int
	Duration int
}

func main() {

	var numberRegex = regexp.MustCompile(`\d+`)
	var boatraces = make([]Boatrace, 0)

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
	scanner.Scan()
	// Time
	timeLines := numberRegex.FindAllString(scanner.Text(), -1)

	scanner.Scan()
	// Duration
	durations := numberRegex.FindAllString(scanner.Text(), -1)

	if len(timeLines) == len(durations) {
		for i := 0; i < len(timeLines); i++ {
			convertedTime, _ := strconv.Atoi(timeLines[i])
			convertedDuration, _ := strconv.Atoi(durations[i])
			boatraces = append(boatraces, Boatrace{
				Time:     convertedTime,
				Duration: convertedDuration,
			})
		}
	}
	
	var result = 1
	for _, race := range boatraces {
		result *= holdButton(race)
	}
	
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