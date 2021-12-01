package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}
func main() {
	var prevSum int
	var currSum int
	var oldNum int
	var newNum int
	var window []int
	var incr int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := 0; i < 3; i++ {
		sc.Scan()
		num, err := strconv.Atoi(sc.Text())
		window = append(window, num)
		if err != nil {
			log.Fatal(err)
		}
	}
	currSum = sum(window)
	for sc.Scan() {
		oldNum = window[0]
		newNum, err = strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal(err)
		}
		prevSum = currSum
		currSum = prevSum - oldNum + newNum
		if currSum > prevSum {
			incr++
		}
		window = append(window[1:], newNum)
	}
	fmt.Printf("increases: %d\n", incr)
}
