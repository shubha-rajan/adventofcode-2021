package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scoreMap := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	pairMap := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	var score int
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var stack []rune
		for _, r := range sc.Text() {
			if pairMap[r] == 0 {
				stack = append(stack, r)
			} else {
				pop := stack[len(stack)-1]
				if !(pairMap[r] == pop) {
					score += scoreMap[r]
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}

	}

	fmt.Printf("result: %d \n", score)
}
