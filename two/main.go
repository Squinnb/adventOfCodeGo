package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("'I love Mondays, back on work site, no more nagging wife.' - Garfield 😤 ")
	txtF, err := ioutil.ReadFile("rps.txt")
	if err != nil {
		log.Fatal(err)
	}
	partOne(string(txtF))
	partTwo(string(txtF))
}
func partTwo(txtFi string) {
	// a rock, b paper, c scis ~ Opponent
	// x rock, y paper, z scis ~ Me
	// X means you need to lose, Y means you need to draw, and Z means you need to win
	// Win Lose Score : (0 if you lost, 3 if the round was a draw, and 6 if you won)
	var score int = 0
	var rpsPointMap = map[string]int{"X": 1, "Y": 2, "Z": 3}
	var tieMap = map[string]string{"A": "X", "B": "Y", "C": "Z"}
	var winMap = map[string]string{"A": "Y", "B": "Z", "C": "X"}
	var loseMap = map[string]string{"A": "Z", "B": "X", "C": "Y"}
	rounds := strings.Split(txtFi, "\n")
	for _, r := range rounds {
		oppMe := strings.Split(r, " ")
		opp := oppMe[0]
		me := oppMe[1]
		if me == "X" {
			// i lose
			score += rpsPointMap[loseMap[opp]]
			score += 0
		} else if me == "Y" {
			// i draw/tie
			score += rpsPointMap[tieMap[opp]]
			score += 3
		} else if me == "Z" {
			// i win
			score += rpsPointMap[winMap[opp]]
			score += 6
		}
	}
	fmt.Println(score)
}
func partOne(txtFi string) {
	// a rock, b paper, c scis ~ Opponent
	// x rock, y paper, z scis ~ Me
	// Shape Score: (1 for Rock, 2 for Paper, and 3 for Scissors)
	// Win Lose Score : (0 if you lost, 3 if the round was a draw, and 6 if you won)
	var score int = 0
	var rpsPointMap = map[string]int{"X": 1, "Y": 2, "Z": 3}
	var rpsMap = map[string]string{"X": "A", "Y": "B", "Z": "C"}
	rounds := strings.Split(txtFi, "\n")
	for _, r := range rounds {
		oppMe := strings.Split(r, " ")
		opp := oppMe[0]
		me := oppMe[1]
		score += rpsPointMap[me]
		if rpsMap[me] == opp {
			score += 3
		} else if me == "X" {
			if opp == "B" {
				score += 0
			} else if opp == "C" {
				score += 6
			}
		} else if me == "Y" {
			if opp == "A" {
				score += 6
			} else if opp == "C" {
				score += 0
			}
		} else if me == "Z" {
			if opp == "A" {
				score += 0
			} else if opp == "B" {
				score += 6
			}
		}
	}
	fmt.Println(score)
}
