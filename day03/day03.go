package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func part1(rucksacks [][]string) int {
	sum := 0
	for _, rucksack := range rucksacks {
		fmt.Println(rucksack)
		var shared string
		for _, char := range rucksack[0] {
			i := strings.Index(rucksack[1], string(char))
			if i > -1 {
				shared = string(char)
				break
			}
		}
		sum += strings.Index(letters, shared) + 1
	}
	return sum
}

func part2(rucksacks []string) int {
	sum := 0

	for i := 0; i < len(rucksacks); i += 3 {
		var shared string
		for _, char := range rucksacks[i] {
			n := strings.Index(rucksacks[i+1], string(char))
			if n > -1 {
				m := strings.Index(rucksacks[i+2], string(char))
				if m > -1 {
					shared = string(char)
					break
				}
			}
		}
		sum += strings.Index(letters, shared) + 1
	}

	return sum
}

func main() {
	file, err := os.Open("day03_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var rucksacks [][]string
	var rucksacks2 []string
	for scanner.Scan() {
		line := scanner.Text()
		rucksacks2 = append(rucksacks2, line)
		rucksack := [2]string{line[0 : len(line)/2], line[len(line)/2 : len(line)]}
		rucksacks = append(rucksacks, rucksack[:])
	}

	fmt.Println(part1(rucksacks))
	fmt.Println(part2(rucksacks2))
}
