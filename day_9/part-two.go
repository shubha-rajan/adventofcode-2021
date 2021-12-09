package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Point struct {
	i int
	j int
}

// Helper to find min in slice of ints
func min(ints []int) (m int) {
	if len(ints) > 0 {
		m = ints[0]
	}
	for i := 1; i < len(ints); i++ {
		if ints[i] < m {
			m = ints[i]
		}
	}
	return m
}

// Helper to determine if current coordinates are a low point
func isLowPoint(i int, j int, grid [][]int) bool {
	var neighbors []int
	curr := grid[i][j]
	if i > 0 {
		neighbors = append(neighbors, grid[i-1][j])
	}
	if j > 0 {
		neighbors = append(neighbors, grid[i][j-1])
	}
	if i < len(grid)-1 {
		neighbors = append(neighbors, grid[i+1][j])
	}
	if j < len(grid[i])-1 {
		neighbors = append(neighbors, grid[i][j+1])
	}
	if curr < min(neighbors) {
		return true
	}
	return false
}

func main() {
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
	var basinSizes []int
	visited := make(map[Point]bool)
	var neighbors []Point
	var curr Point

	for i, row := range grid {
		for j, _ := range row {
			curr = Point{i: i, j: j}
			if !visited[curr] && isLowPoint(i, j, grid) {
				neighbors = append(neighbors, curr)
				basinSize := 0
				// it's BFS time
				for len(neighbors) > 0 {
					curr, neighbors = neighbors[0], neighbors[1:]
					currVal := grid[curr.i][curr.j]
					if !visited[curr] && currVal != 9 {
						basinSize += 1
						visited[curr] = true
						if curr.i > 0 && grid[curr.i-1][curr.j] >= currVal {
							neighbors = append(neighbors, Point{i: curr.i - 1, j: curr.j})
						}
						if curr.j > 0 && grid[curr.i][curr.j-1] >= currVal {
							neighbors = append(neighbors, Point{i: curr.i, j: curr.j - 1})
						}
						if curr.i < len(grid)-1 && grid[curr.i+1][curr.j] >= currVal {
							neighbors = append(neighbors, Point{i: curr.i + 1, j: curr.j})
						}
						if curr.j < len(row)-1 && grid[curr.i][curr.j+1] >= currVal {
							neighbors = append(neighbors, Point{i: curr.i, j: curr.j + 1})
						}
					}
				}
				basinSizes = append(basinSizes, basinSize)
			}
		}

	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	product := 1
	for _, size := range basinSizes[:3] {
		product *= size
	}

	fmt.Printf("basinSizes: %d\n", basinSizes)
	fmt.Printf("product: %d\n", product)
}
