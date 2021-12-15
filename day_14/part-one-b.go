package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Node struct {
	val  string
	next *Node
}

func main() {
	var head *Node
	rules := make(map[string]string)
	counts := make(map[string]int)
	steps := 10
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	// build the linked list
	var curr *Node
	for _, r := range sc.Text() {
		counts[string(r)]++
		if head == nil {
			head = &Node{val: string(r), next: nil}
			curr = head
		} else {
			curr.next = &Node{val: string(r), next: nil}
			curr = curr.next
		}
	}
	sc.Scan() // Skip the blank line
	for sc.Scan() {
		// build the map
		rule := strings.Split(sc.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}
	for i := 0; i < steps; i++ {
		curr = head
		for curr.next != nil {
			pair := curr.val + curr.next.val
			tmp := curr.next
			curr.next = &Node{val: rules[pair], next: tmp}
			curr = tmp
			counts[rules[pair]]++
		}
	}
	minCt := math.MaxInt64
	var maxCt int
	for _, ct := range counts {
		if ct > maxCt {
			maxCt = ct
		} else if ct < minCt {
			minCt = ct
		}
	}
	fmt.Printf("count: %v \n", counts)
	fmt.Printf("result: %d \n", maxCt-minCt)
}
