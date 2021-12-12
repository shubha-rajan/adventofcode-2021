package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func loadMap(sc *bufio.Scanner) map[string][]string {
	caveMap := make(map[string][]string)
	for sc.Scan() {
		caves := strings.Split(sc.Text(), "-")
		caveMap[caves[0]] = append(caveMap[caves[0]], caves[1])
		caveMap[caves[1]] = append(caveMap[caves[1]], caves[0])
	}
	return caveMap
}

func countPaths(cave string, caveMap *map[string][]string, pathCount *int, visited *map[string]bool) {
	if cave == "end" {
		*pathCount += 1
		return
	}
	for _, neighbor := range (*caveMap)[cave] {
		if !(*visited)[neighbor] {
			if strings.ToLower(neighbor) == neighbor {
				(*visited)[neighbor] = true
			}
			countPaths(neighbor, caveMap, pathCount, visited)
			if strings.ToLower(neighbor) == neighbor {
				(*visited)[neighbor] = false
			}
		}
	}

}
func main() {
	visited := make(map[string]bool)
	var pathCount int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	caveMap := loadMap(bufio.NewScanner(file))

	visited["start"] = true
	countPaths("start", &caveMap, &pathCount, &visited)

	fmt.Printf("result: %d \n", pathCount)

}
