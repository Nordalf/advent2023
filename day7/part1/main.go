package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var cardValueMap = map[string]int {
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type Hand struct {
	Cards []string
	Bid int
	Type cardType
	Points int
	Rank int
}

type cardType string
const (
	FiveOf cardType = "6"
	FourOf cardType = "5"
	FullHouse cardType = "4"
	ThreeOf cardType = "3"
	TwoPair cardType = "2"
	OnePair cardType = "1"
	HighCard cardType = "0"
)

type ByType []Hand

func (a ByType) Len() int           { return len(a) }
func (a ByType) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByType) Less(i, j int) bool { return a[i].Type > a[j].Type}

type ByCardType []Hand

func (a ByCardType) Len() int { return len(a) }
func (a ByCardType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCardType) Less(i, j int) bool {
	for it := 0; it < 5; it++ {
		log.Println(a[i].Cards, i, a[j].Cards, j)
		if cardValueMap[a[i].Cards[it]] == cardValueMap[a[j].Cards[it]] && a[i].Type == a[j].Type {
			continue
		}
		if cardValueMap[a[i].Cards[it]] > cardValueMap[a[j].Cards[it]] && a[i].Type == a[j].Type {
//			log.Println("it worked: ", a[i].Cards, a[j].Cards)
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	var hands []Hand
	var sum int
	handsFile, err := os.Open("../../datasets/hands.txt")
	if err != nil {
		log.Printf("error occurred reading file: %v", err.Error())
		return
	}

	defer func(handsFile *os.File) {
		closeErr := handsFile.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}(handsFile)

	scanner := bufio.NewScanner(handsFile)
	var line []string
	var re = regexp.MustCompile(`\w`)
	var bid int
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), " ")
		if len(line) < 2 {
			log.Fatal("the line did not contain the hand a bid")
		}
		bid, _ = strconv.Atoi(line[1])
		hands = append(hands, Hand{
			Cards:  re.FindAllString(line[0], -1),
			Bid:    bid,
			Type:   "",
			Points: 0,
		})
	}
	
	for i, _ := range hands {
		hands[i].FindType()
	}
	
	sort.Sort(ByType(hands))
//	for _, hand := range hands {
//		log.Println(hand.Cards, hand.Type)
//	}
	log.Println()
	sort.Sort(ByCardType(hands))
//	for _, hand := range hands {
//		log.Println(hand.Cards, hand.Type)
//	}

	// Calculate the total sum
	for i, hand := range hands {
		sum += hand.Bid * (len(hands)-i)
	}
	
	log.Println(sum)
}

func (h *Hand) FindType() {
	var cardMap = make(map[string]int)
	
	for _, card := range h.Cards {
		if _, ok := cardMap[card]; ok {
			cardMap[card] += 1
		} else {
			cardMap[card] = 1
		}
		h.Points += cardValueMap[card]
	}
	
	
	// If this is the case, it means, all the cards are distinct
	if len(cardMap) == 5 {
		h.Type = HighCard
	} else {
		pair := 0
		
		for _, v := range cardMap {
			switch v {
			case 5:
				h.Type = FiveOf
			case 4:
				h.Type = FourOf
			case 3:
				h.Type = ThreeOf
				pair++
			case 2:
				h.Type = OnePair
				pair++
			}
		}
		
		if pair == 2 && len(cardMap) == 2 {
			h.Type = FullHouse
		}
		
		if pair == 2 && len(cardMap) == 3 {
			h.Type = TwoPair
		}
	}
}
