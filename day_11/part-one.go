package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


type Point struct {
	i int
	j int
}

func loadGrid(sc *bufio.Scanner) (grid [10][10]int) {
	var i int
	for sc.Scan() {
		for j, r := range sc.Text() {
			grid[i][j] = int(r - '0')
		}
		i++
	}
	return grid
}

func getNeighbors(p Point) (neighbors []Point) {

	if p.i > 0 {
		neighbors = append(neighbors, Point{i: p.i - 1, j: p.j})
	}
	if p.j > 0 {
		neighbors = append(neighbors, Point{i: p.i, j: p.j - 1})
	}
	if p.i < 9 {
		neighbors = append(neighbors, Point{i: p.i + 1, j: p.j})
	}
	if p.j < 9 {
		neighbors = append(neighbors, Point{i: p.i, j: p.j + 1})
	}
	if p.i > 0 && p.j > 0 {
		neighbors = append(neighbors, Point{i: p.i - 1, j: p.j - 1})
	}
	if p.i > 0 && p.j < 9 {
		neighbors = append(neighbors, Point{i: p.i - 1, j: p.j + 1})
	}
	if p.i < 9 && p.j > 0 {
		neighbors = append(neighbors, Point{i: p.i + 1, j: p.j - 1})
	}
	if p.i < 9 && p.j < 9 {
		neighbors = append(neighbors, Point{i: p.i + 1, j: p.j + 1})
	}
	return neighbors
}

func flash(grid *[10][10]int) (flashes int) {
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] > 9 {
				neighbors := getNeighbors(Point{i: i, j: j})
				flashes += 1
				for _, p := range neighbors {
					grid[p.i][p.j] += 1
				}
				grid[i][j] = -10
			}
		}
	}
	if flashes > 0 {
		return flashes + flash(grid)
	} else {
		return flashes
	}
}

func main() {
	var grid [10][10]int
	var flashes int
	steps := 100
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	grid = loadGrid(bufio.NewScanner(file))

	for count := 0; count < steps; count++ {
		for i, row := range grid {
			for j, _ := range row {
				grid[i][j] += 1
			}
		}
		flashes += flash(&grid)

		for i, row := range grid {
			for j, _ := range row {
				if grid[i][j] < 0 {
					grid[i][j] = 0
				}
			}
		}

	}
	fmt.Printf("result: %d \n", flashes)

}
