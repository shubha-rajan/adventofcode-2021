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
	var sum int
	var nums []string
	var target byte

	var oxygen []string
	var scrubber []string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		nums = append(nums, sc.Text())
	}

	var tmpOxy []string
	oxygen = nums
	i := 0
	for len(oxygen) > 1 {
		sum = 0
		for _, numStr := range oxygen {
			if numStr[i] == '1' {
				sum += 1
			}
		}

		if sum*2 >= len(oxygen) {
			target = '1'
		} else {
			target = '0'
		}

		for _, numStr := range oxygen {
			if numStr[i] == target {
				tmpOxy = append(tmpOxy, numStr)
			}
		}
		oxygen = tmpOxy
		tmpOxy = nil
		i++
	}

	var tmpScrub []string
	scrubber = nums
	i = 0
	for len(scrubber) > 1 {
		sum = 0
		for _, numStr := range scrubber {
			if numStr[i] == '1' {
				sum += 1
			}
		}
		if sum*2 >= len(scrubber) {
			target = '0'
		} else {
			target = '1'
		}
		for _, numStr := range scrubber {
			if numStr[i] == target {
				tmpScrub = append(tmpScrub, numStr)
			}
		}

		scrubber = tmpScrub
		tmpScrub = nil
		i++

	}

	oxy, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	scrub, err := strconv.ParseInt(scrubber[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// print result here
	fmt.Printf("Oxygen: %v\n", oxygen[0])
	fmt.Printf("Scrubber: %v\n", scrubber[0])
	fmt.Printf("result: %d\n", oxy*scrub)

}
