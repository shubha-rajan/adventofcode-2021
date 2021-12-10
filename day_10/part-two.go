package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	scoreMap := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	pairMap := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	var scores []int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var stack []rune
		score := 0
		isCorrupted := false
		for _, r := range sc.Text() {
			if pairMap[r] == 0 {
				stack = append(stack, r)
			} else {
				pop := stack[len(stack)-1]
				if !(pairMap[r] == pop) {
					isCorrupted = true
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if !isCorrupted {
			for i := len(stack) - 1; i >= 0; i-- {
				score *= 5
				score += scoreMap[stack[i]]
			}
			scores = append(scores, score)
		}

	}
	sort.Sort(sort.IntSlice(scores))

	fmt.Printf("result: %d \n", scores)
	fmt.Printf("result: %d \n", scores[len(scores)/2])

}
