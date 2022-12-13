package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func removeItem(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}

func popElem(stack []string, number int) ([]string, []string) {
	size := len(stack)
	return stack[:size-number], stack[size-number:]
}

func part1(stacks [][]string, commands [][]string) string {

	for _, command := range commands {
		count, _ := strconv.Atoi(command[0])
		from, _ := strconv.Atoi(command[1])
		to, _ := strconv.Atoi(command[2])
		for i := 0; i < count; i++ {
			elem := []string{}

			stacks[from-1], elem = popElem(stacks[from-1], 1)
			stacks[to-1] = append(stacks[to-1], elem...)
		}
	}

	retVal := ""
	for _, stack := range stacks {
		retVal += stack[len(stack)-1]
	}
	return retVal
}

func part2(stacks [][]string, commands [][]string) string {

	for _, command := range commands {
		count, _ := strconv.Atoi(command[0])
		from, _ := strconv.Atoi(command[1])
		to, _ := strconv.Atoi(command[2])
		elem := []string{}
		stacks[from-1], elem = popElem(stacks[from-1], count)
		stacks[to-1] = append(stacks[to-1], elem...)
	}

	retVal := ""
	for _, stack := range stacks {
		retVal += stack[len(stack)-1]
	}
	return retVal
}

func main() {
	file, err := os.Open("day05_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	stacks := [][]string{
		{"D", "T", "W", "F", "J", "S", "H", "N"},
		{"H", "R", "P", "Q", "T", "N", "B", "G"},
		{"L", "Q", "V"},
		{"N", "B", "S", "W", "R", "Q"},
		{"N", "D", "F", "T", "V", "M", "B"},
		{"M", "D", "B", "V", "H", "T", "R"},
		{"D", "B", "Q", "J"},
		{"D", "N", "J", "V", "R", "Z", "H", "Q"},
		{"B", "N", "H", "M", "S"},
	}

	var commands [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		lineSplit = removeItem(lineSplit, 0)
		lineSplit = removeItem(lineSplit, 1)
		lineSplit = removeItem(lineSplit, 2)
		commands = append(commands, lineSplit)
	}
	copy1 := make([][]string, len(stacks))
	copy2 := make([][]string, len(stacks))
	for i := range stacks {
		copy1[i] = make([]string, len(stacks[i]))
		copy2[i] = make([]string, len(stacks[i]))
		copy(copy1[i], stacks[i])
		copy(copy2[i], stacks[i])
	}
	fmt.Println(part1(copy1, commands))
	fmt.Println(part2(copy2, commands))
}
