package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func detectDuplicates(slice []string) bool {
	fmt.Println(slice)
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return len(list) != 14
}

func part1(datastream string) int {
	datastreamSlice := strings.Split(datastream, "")
	for i := 0; i < len(datastreamSlice)-14; i++ {
		if !detectDuplicates(datastreamSlice[i : i+14]) {
			return i + 14
		}
	}
	return -1
}

func main() {
	file, err := ioutil.ReadFile("day06_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	datastream := string(file)
	fmt.Println(part1(datastream))
}
