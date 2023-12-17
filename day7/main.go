package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var (
	strength1 = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"} // cards relative strength
	strength2 = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
)

type Hand struct {
	Cards string // KK677
	Bid   int    // 100
	Not   []int  // Camel Cards strength notation, e.g [1, 1, 3] for three of a kind
}

type Hands []Hand

/*
	Camel Cards strength notation in descending order

	five of a kind [5]
	four of a kind [1, 4]
	full house [2, 3]
	three of a kind [1, 1, 3]
	two pairs [1, 2, 2]
	one pair [1, 1, 1, 2]
	high card [1, 1, 1, 1, 1]
*/

/*
	Part Two: J is considered a Joker instead of a Jack
	If one or serveral Jokers are present, the hand is upgraded to the highest possible hand:

	five of a kind [5] 			=> no effect
	four of a kind [1, 4] 		=> upgrate to [5]
	full house [2, 3] 			=> upgrade to [5] (JJxxx, JJJxx)
	three of a kind [1, 1, 3] 	=> upgrade to [1,4] (Jxxx, JJJx)
	two pairs [1, 2, 2] 		=> upgrade to [2,3] or [1,4] (xxJyy, JJxx)
	one pair [1, 1, 1, 2] 		=> upgrade to [1,1,3] (Jxx, JJx)
	high card [1, 1, 1, 1, 1] 	=> upgrade to [1,1,1,2] (Jx)
*/

func main() {
	hands := readFile("input.txt")
	classifyHands(&hands)

	// Part One
	sortHands(&strength1, &hands)
	fmt.Println("Part One: ", totalWinnings(&hands))

	// Part Two
	upgradeJokers(&hands)
	sortHands(&strength2, &hands)
	fmt.Println("Part Two: ", totalWinnings(&hands))
}

// Classify each hand according to Camel Cards strength notation
func classifyHands(hands *Hands) {
	for i, hand := range *hands {
		hash := make(map[string]int) // hash map of cards count
		slice := []int{}             // converted to slice of counts, sorted [1, 1, 3]

		for _, card := range hand.Cards {
			if _, ok := hash[string(card)]; ok {
				hash[string(card)]++
			} else {
				hash[string(card)] = 1
			}
		}

		for _, v := range hash {
			slice = append(slice, v)
			sort.Ints(slice)
		}
		(*hands)[i].Not = slice
	}
}

// Upgrate Joker hands 
func upgradeJokers(hands *Hands) {
	for i, hand := range *hands {
		numJokers := strings.Count(hand.Cards, "J")

		if numJokers == 0 {
			continue
		}

		currentEval := hand.Not
		currentEvalLen := len(currentEval)

		n := []int{5}

		switch currentEvalLen {
		case 3:
			if currentEval[2] == 3 { // three of a kind
				n = []int{1, 4}
			} else if currentEval[2] == 2 { // two pairs
				if numJokers == 1 {
					n = []int{2, 3}
				} else if numJokers == 2 {
					n = []int{1, 4}
				}
			}
		case 4: // one pair
			if numJokers == 1 || numJokers == 2 {
				n = []int{1, 1, 3}
			}
		case 5: // high card
			if numJokers == 1 {
				n = []int{1, 1, 1, 2}
			}
		}
		(*hands)[i].Not = n
	}
}

// Sort hands according to Camel Cards strength notation
func sortHands(s *[]string, hands *Hands) {
	slices.SortStableFunc(*hands, func(a, b Hand) int {
		lenA := len(a.Not)
		lenB := len(b.Not)

		if lenA > lenB {
			return -1
		} else if lenA < lenB {
			return 1
		}

		lastA := a.Not[lenA-1]
		lastB := b.Not[lenB-1]

		// Now same length, then compare last element
		if lastA > lastB {
			return 1
		} else if lastA < lastB {
			return -1
		}

		// Else, return highest card
		for i := 0; i < 5; i++ {
			ai := slices.Index(*s, string(a.Cards[i]))
			bi := slices.Index(*s, string(b.Cards[i]))

			if ai == bi {
				continue
			} else if ai > bi {
				return 1
			} else if ai < bi {
				return -1
			}
		}
		return 0 
	})
}

func totalWinnings(hands *Hands) int {
	totalWinnings := 0
	for i, hand := range *hands {
		rank := i + 1
		totalWinnings += hand.Bid * rank
	}
	return totalWinnings
}

func readFile(fileName string) Hands {
	var hands Hands

	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		bid, err := strconv.Atoi(parts[1])
		check(err)

		hand := Hand{
			Cards: parts[0],
			Bid:   bid,
			Not:   []int{},
		}
		hands = append(hands, hand)
	}
	return hands
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
