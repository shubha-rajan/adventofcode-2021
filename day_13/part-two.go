package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func stringToPoint(s string) (Point, int, int) {
	coords := strings.Split(s, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{x: x, y: y}, x, y
}

func fold(
	step string, 
	points *map[Point]bool, 
	maxX int, 
	maxY int) (axis string, value int) {
	fold := strings.Split(strings.Fields(step)[2], "=")
	value, err := strconv.Atoi(fold[1])
	if err != nil {
		log.Fatal(err)
	}
	axis = fold[0]
	if axis == "x" {
		x := value
		foldSize := maxX - x
		for y := 0; y <= maxY; y++ {
			for i := 1; i <= foldSize; i++ {
				p := Point{x: x + i, y: y}
				if (*points)[p] {
					delete(*points, p)
					(*points)[Point{x: x - i, y: y}] = true
				}
			}
		}
	} else if axis == "y" {
		y := value
		foldSize := maxY - y
		for x := 0; x <= maxX; x++ {
			for i := 1; i <= foldSize; i++ {
				p := Point{x: x, y: y + i}
				if (*points)[p] {
					delete(*points, p)
					(*points)[Point{x: x, y: y - i}] = true
				}
			}
		}
	}
	return axis, value
}

func main() {
	points := make(map[Point]bool)
	var maxX int
	var maxY int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	for sc.Text() != "" {
		p, x, y := stringToPoint(sc.Text())
		points[p] = true
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		sc.Scan()
	}

	for sc.Scan() {
		axis, value := fold(sc.Text(), &points, maxX, maxY)
		if axis == "x" {
			maxX = value
		}
		if axis == "y" {
			maxY = value
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if points[Point{x: x, y: y}] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	fmt.Printf("points: %v \n", len(points))
}
