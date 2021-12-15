package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	var polymer string
	rules := make(map[string]string)
	counts := make(map[string]int)
	pairCounts := make(map[string]int)
	steps := 40
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	polymer = sc.Text()
	for i, r := range polymer {
		counts[string(r)]++
		if i < len(polymer)-1 {
			pairCounts[polymer[i:i+2]]++
		}
	}
	sc.Scan() // Skip the blank line
	for sc.Scan() {
		// build the map
		rule := strings.Split(sc.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}

	for i := 0; i < steps; i++ {
		fmt.Printf("pairCounts: %v,  \n", pairCounts)
		newPairCounts := make(map[string]int)
		for pair, count := range pairCounts {
			counts[rules[pair]] += count
			newPairCounts[string(pair[0])+rules[pair]] += count
			newPairCounts[rules[pair]+string(pair[1])] += count
		}
		pairCounts = newPairCounts
	}

	minCt := math.MaxInt64
	var maxCt int
	for _, ct := range counts {
		if ct > maxCt {
			maxCt = ct
		} else if ct < minCt {
			minCt = ct
		}
	}
	fmt.Printf("result: %v  \n", maxCt-minCt)
}
