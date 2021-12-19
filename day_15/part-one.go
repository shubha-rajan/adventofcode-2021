package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadGrid(sc *bufio.Scanner) (grid [][]int) {
	for sc.Scan() {
		var row []int
		for _, r := range sc.Text() {
			row = append(row, (int(r - '0')))
		}
		grid = append(grid, row)
	}
	return grid
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	var grid [][]int
	var riskMatrix [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	grid = loadGrid(bufio.NewScanner(file))

	for i := 0; i < len(grid); i++ {
		riskMatrix = append(riskMatrix, make([]int, len(grid[0])))
	}

	riskMatrix[0][0] = grid[0][0]

	for i := 1; i < len(grid); i++ {
		riskMatrix[i][0] = riskMatrix[i-1][0] + grid[i][0]
	}

	for j := 1; j < len(grid[0]); j++ {
		riskMatrix[0][j] = riskMatrix[0][j-1] + grid[0][j]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			riskMatrix[i][j] = min(riskMatrix[i-1][j], riskMatrix[i][j-1]) + grid[i][j]
		}
	}

	last := len(grid) - 1
	fmt.Printf("result: %d \n", riskMatrix[last][last]-grid[0][0])

}
