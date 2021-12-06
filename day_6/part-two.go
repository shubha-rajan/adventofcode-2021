package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	var ages [9]int
	days := 256
	var popSize int
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sliceData := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), ",")
	for _, s := range sliceData {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ages[n] += 1
	}

	for i := 0; i < days; i++ {
		zeros := ages[0]
		for j := 0; j < 8; j++ {
			ages[j] = ages[j+1]
		}
		ages[8] = zeros
		ages[6] += zeros
	}
	for _, count := range ages {
		popSize += count
	}

	fmt.Printf("ages: %v\n", ages)
	fmt.Printf("fish: %d\n", popSize)
}
