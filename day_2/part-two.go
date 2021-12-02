package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var depth int
	var pos int
	var aim int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		move := strings.Fields(sc.Text())
		dist, err := strconv.Atoi(move[1])
		if err != nil {
			log.Fatal(err)
		}
		switch move[0] {
		case "forward":
			pos += dist
			depth += (aim * dist)
		case "down":
			aim += dist
		case "up":
			aim -= dist
		}
	}
	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("posn: %d\n", pos)
	fmt.Printf("product: %d\n", pos*depth)
}
