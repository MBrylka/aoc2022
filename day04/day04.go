package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(pairs []string) int {
	sum := 0
	for _, pair := range pairs {
		elfs := strings.Split(pair, ",")
		first := strings.Split(elfs[0], "-")
		second := strings.Split(elfs[1], "-")

		l1, _ := strconv.Atoi(string(first[0]))
		p1, _ := strconv.Atoi(string(first[1]))
		l2, _ := strconv.Atoi(string(second[0]))
		p2, _ := strconv.Atoi(string(second[1]))
		if (l1 >= l2 && p1 <= p2) || (l2 >= l1 && p2 <= p1) {
			fmt.Println(elfs)
			sum += 1
		}
	}

	return sum
}

func beetween(lx int, l1 int, l2 int) bool {
	return lx >= l1 && lx <= l2
}

func part2(pairs []string) int {
	sum := 0
	for _, pair := range pairs {
		elfs := strings.Split(pair, ",")
		first := strings.Split(elfs[0], "-")
		second := strings.Split(elfs[1], "-")

		l1, _ := strconv.Atoi(string(first[0]))
		p1, _ := strconv.Atoi(string(first[1]))
		l2, _ := strconv.Atoi(string(second[0]))
		p2, _ := strconv.Atoi(string(second[1]))

		if beetween(l1, l2, p2) ||
			beetween(p1, l2, p2) ||
			beetween(l2, l1, p1) ||
			beetween(p2, l1, p1) {
			fmt.Println(elfs)
			sum += 1
		}
	}

	return sum
}

func main() {
	file, err := os.Open("day04_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var pairs []string

	for scanner.Scan() {
		line := scanner.Text()
		pairs = append(pairs, line)
	}

	//fmt.Println(part1(pairs))
	fmt.Println(part2(pairs))
}
