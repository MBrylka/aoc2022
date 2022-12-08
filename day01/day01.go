package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day01_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var l [][]int
	var l2 []int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			intVar, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			l2 = append(l2, intVar)
		} else {
			l = append(l, l2)
			l2 = nil
		}
	}
	var max []int
	for _, elf := range l {
		sum := 0
		for _, calorie := range elf {
			sum += calorie
		}
		max = append(max, sum)
	}

	sort.Ints(max)
	fmt.Println(max[len(max)-1])
	fmt.Println(max[len(max)-2])
	fmt.Println(max[len(max)-3])
}
