package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func min(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
		}
	}
	return m
}
func main() {
	var riskLevelSum int
	var grid [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var row []int
		for _, r := range sc.Text() {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}

	var neighbors []int
	var curr int
	for i, row := range grid {
		for j, val := range row {
			curr = val
			neighbors = nil
			if i > 0 {
				neighbors = append(neighbors, grid[i-1][j])
			}
			if j > 0 {
				neighbors = append(neighbors, grid[i][j-1])
			}
			if i < len(grid)-1 {
				neighbors = append(neighbors, grid[i+1][j])
			}
			if j < len(row)-1 {
				neighbors = append(neighbors, grid[i][j+1])
			}
			if curr < min(neighbors) {
				riskLevelSum += curr + 1
			}
		}

	}
	fmt.Printf("riskLevelSum: %d\n", riskLevelSum)
}
