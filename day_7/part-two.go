package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	positions := make(map[int]int)
	var maxPosition int
	minPosition := math.MaxInt64
	var currFuel int
	minFuel := math.MaxInt64
	var target int
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
		positions[n] += 1
		if n > maxPosition {
			maxPosition = n
		}
		if n < minPosition {
			minPosition = n
		}
	}

	for i := minPosition; i < maxPosition; i++ {
		currFuel = 0
		for posn, ct := range positions {
			distance := abs(posn - i)
			var moveFuel int
			for j := 1; j <= distance; j++ {
				moveFuel += j
			}
			currFuel += ct * moveFuel
		}
		if currFuel < minFuel {
			minFuel = currFuel
			target = i

		}

	}

	fmt.Printf("minFuel: %d\n", minFuel)
	fmt.Printf("target: %d\n", target)
}
