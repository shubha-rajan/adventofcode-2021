package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var curr int
	var prev int
	var incr int
	first := true
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		prev = curr
		curr, err = strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal(err)
		}
		if !first && curr > prev {
			incr++
		}
		if first {
			first = false
		}
	}
	fmt.Printf("increases: %d\n", incr)
}
