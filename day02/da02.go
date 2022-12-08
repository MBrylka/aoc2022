package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//ROCK PAPER SCISSORS
// A     B     C
// X     Y     Z

func part1(games []string) int {
	gameMap := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}
	sum := 0
	for _, game := range games {
		//fmt.Println(game, gameMap[game])
		sum += gameMap[game]
	}
	return sum
}

func part2(games []string) int {
	// X     Y     Z
	// win  draw  loose

	//ROCK PAPER SCISSORS
	// A     B     C
	gameMap := map[string]string{
		"A X": "C",
		"A Y": "A",
		"A Z": "B",
		"B X": "A",
		"B Y": "B",
		"B Z": "C",
		"C X": "B",
		"C Y": "C",
		"C Z": "A",
	}

	scoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	sum := 0
	for _, game := range games {
		sum += scoreMap[gameMap[game]] //add score based on figure used
		sum += scoreMap[game[2:3]]
	}
	return sum
}

func main() {
	file, err := os.Open("day02_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var games []string

	for scanner.Scan() {
		games = append(games, scanner.Text())
	}

	fmt.Println(part1(games))
	fmt.Println(part2(games))
}
