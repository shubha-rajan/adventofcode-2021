package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	var fishes []int
	days := 10
	fileBytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	sliceData := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), ",")
	for _, s := range sliceData {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fishes = append(fishes, n)
	}

	for i := 0; i < days; i++ {
		for j := 0; j < len(fishes); j++ {
			if fishes[j] == 0 {
				fishes[j] = 6
				// set the new fish to 9 because it will decrement to 8 by the end of this iteration
				fishes = append(fishes, 9)
			} else {
				fishes[j] -= 1
			}
		}

	}

	fmt.Printf("fish: %d\n", len(fishes))
}
