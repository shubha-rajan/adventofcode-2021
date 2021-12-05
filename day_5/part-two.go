package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// helper functions
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// helpful structs
type Point struct {
	x int
	y int
}

type Line struct {
	start *Point
	end   *Point
}

func stringToPoint(s string) Point {
	coords := strings.Split(s, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal(err)
	}

	return Point{x: x, y: y}

}
func (l *Line) getAllPoints() []Point {
	var points []Point

	if l.start.x == l.end.x {
		maxY := max(l.start.y, l.end.y)
		minY := min(l.start.y, l.end.y)
		for i := minY; i <= maxY; i++ {
			points = append(points, Point{x: l.start.x, y: i})
		}
	} else if l.start.y == l.end.y {
		maxX := max(l.start.x, l.end.x)
		minX := min(l.start.x, l.end.x)
		for i := minX; i <= maxX; i++ {
			points = append(points, Point{x: i, y: l.start.y})
		}
	} else {
		slope := (l.end.y - l.start.y) / (l.end.x - l.start.x)
		xDist := abs(l.end.x - l.start.x)
		if slope < 0 && l.end.x > l.start.x {
			for i := 0; i <= xDist; i++ {
				points = append(points, Point{x: l.start.x + i, y: l.start.y - i})
			}
		} else if slope < 0 && l.end.x < l.start.x {
			for i := 0; i <= xDist; i++ {
				points = append(points, Point{x: l.start.x - i, y: l.start.y + i})
			}
		} else if slope > 0 && l.end.x > l.start.x {
			for i := 0; i <= xDist; i++ {
				points = append(points, Point{x: l.start.x + i, y: l.start.y + i})
			}
		} else if slope > 0 && l.end.x < l.start.x {
			for i := 0; i <= xDist; i++ {
				points = append(points, Point{x: l.start.x - i, y: l.start.y - i})
			}
		}
	}
	return points
}

func main() {
	points := make(map[Point]int)
	var count int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ptStrs := strings.Split(sc.Text(), " -> ")
		ptStart := stringToPoint(ptStrs[0])
		ptEnd := stringToPoint(ptStrs[1])

		line := Line{start: &ptStart, end: &ptEnd}
		linePts := line.getAllPoints()
		for _, pt := range linePts {
			points[pt] += 1
		}
	}
	for _, v := range points {
		if v >= 2 {
			count++
		}
	}

	fmt.Printf("number of danger points: %v \n", count)
}
