package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	segCounts := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	var result int
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		row := strings.Split(sc.Text(), " | ")
		outputs := strings.Fields(row[1])

		for _, s := range outputs {
			if len(s) == segCounts[1] ||
				len(s) == segCounts[4] ||
				len(s) == segCounts[7] ||
				len(s) == segCounts[8] {
				result++
			}
		}

	}

	fmt.Printf("result: %d \n", result)
}
