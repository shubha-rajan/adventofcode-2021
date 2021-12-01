package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
		if err != nil {
			log.Fatal(err)
		}
		window = append(window, num)
		currSum += num
	}

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
