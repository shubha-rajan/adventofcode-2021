package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// variables + data structures to track result value
	var sums []int
	var count int
	var gammaStr string
	var epsilonStr string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		for i, c := range sc.Text() {
			if len(sums) == i {
				sums = append(sums, 0)
			}
			if c == '1' {
				sums[i] += 1
			}
		}
		count += 1
	}
	for _, num := range sums {
		if num > (count / 2) {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}
	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// print result here
	fmt.Printf("gamma: %v\n", gammaStr)
	fmt.Printf("epsilon: %v\n", epsilonStr)
	fmt.Printf("result: %d\n", gamma*epsilon)

}
