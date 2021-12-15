package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// alternative recursive solution, for fun
func build(step int, pair string, rules *map[string]string, counts *map[string]int) string {
	if step == 10 {
		return pair
	}

	strings.Split(pair, "")
	(*counts)[(*rules)[pair]]++
	a := build(step+1, string(pair[0])+(*rules)[pair], rules, counts)
	b := build(step+1, (*rules)[pair]+string(pair[1]), rules, counts)

	return a + b[1:]

}
func main() {
	var polymer string
	rules := make(map[string]string)
	counts := make(map[string]int)
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	polymer = sc.Text()
	for _, r := range polymer {
		counts[string(r)]++
	}
	sc.Scan() // Skip the blank line
	for sc.Scan() {
		// build the map
		rule := strings.Split(sc.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}

	result := build(0, polymer[:2], &rules, &counts)
	for i := 1; i < len(polymer)-1; i++ {
		result += build(0, polymer[i:i+2], &rules, &counts)
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
	fmt.Printf("count: %v \n", counts)
	fmt.Printf("result: %v,  \n", maxCt-minCt)
}
