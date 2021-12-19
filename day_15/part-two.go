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

func buildTileMatrix(size int) [][]int {
	var tiles [][]int
	for i := 0; i < size; i++ {
		tiles = append(tiles, make([]int, size))
	}

	for j := 0; j < size; j++ {
		tiles[0][j] = j
	}

	for j := 0; j < size; j++ {
		for i := 1; i < size; i++ {
			tiles[i][j] = tiles[i-1][j] + 1
		}
	}
	return tiles
}

func buildTiledGrid(grid [][]int, tiles [][]int) [][]int {
	var tileGrid [][]int
	for i := 0; i < len(tiles)*len(grid); i++ {
		tileGrid = append(tileGrid, make([]int, len(tiles[0])*len(grid[0])))
		iTile := i / len(grid)
		for j := 0; j < len(tiles[0])*len(grid[0]); j++ {
			jTile := j / len(grid[0])
			risk := grid[i%len(grid)][j%len(grid)] + tiles[iTile][jTile]
			for risk > 9 {
				risk -= 9
			}
			tileGrid[i][j] = risk
		}
	}
	return tileGrid
}
func main() {
	var grid [][]int
	var riskMatrix [][]int
	size := 5
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	tiles := buildTileMatrix(size)
	grid = loadGrid(bufio.NewScanner(file))
	tiledGrid := buildTiledGrid(grid, tiles)

	for i := 0; i < len(grid)*size; i++ {
		riskMatrix = append(riskMatrix, make([]int, len(grid[0])*size))
	}

	riskMatrix[0][0] = tiledGrid[0][0]

	for i := 1; i < len(grid)*size; i++ {
		riskMatrix[i][0] = riskMatrix[i-1][0] + tiledGrid[i][0]
	}

	for j := 1; j < len(grid[0])*size; j++ {
		riskMatrix[0][j] = riskMatrix[0][j-1] + tiledGrid[0][j]
	}

	for i := 1; i < len(grid)*size; i++ {
		for j := 1; j < len(grid[0])*size; j++ {
			riskMatrix[i][j] = min(riskMatrix[i-1][j], riskMatrix[i][j-1]) + tiledGrid[i][j]
		}
	}

	last := len(riskMatrix) - 1
	fmt.Printf("result: %d \n", riskMatrix[last][last]-tiledGrid[0][0])

}
