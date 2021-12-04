package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	state [5][5]string
}

func loadBoards(sc *bufio.Scanner) []*Board {
	var boards []*Board
	var currBoard [5][5]string
	var i int
	for sc.Scan() {
		row := strings.Fields(sc.Text())
		if i != 0 {
			for j, n := range row {
				currBoard[i-1][j] = n
			}
		}
		i = (i + 1) % (6)
		if i == 0 {
			boards = append(boards, &Board{state: currBoard})
			currBoard = [5][5]string{}
		}
	}
	return boards
}

func (b *Board) checkRows(called map[string]bool) bool {
	for i := 0; i < 5; i++ {
		winner := true
		for j := 0; j < 5; j++ {
			if !called[b.state[i][j]] {
				winner = false
			}
		}
		if winner == true {
			return true
		}
	}
	return false
}

func (b *Board) checkColumns(called map[string]bool) bool {
	for j := 0; j < 5; j++ {
		winner := true
		for i := 0; i < 5; i++ {
			if !called[b.state[i][j]] {
				winner = false
			}
		}
		if winner == true {
			return true
		}
	}
	return false
}

func (b *Board) calculateScore(called map[string]bool, current string) int {
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !called[b.state[i][j]] {
				n, err := strconv.Atoi(b.state[i][j])
				if err != nil {
					log.Fatal(err)
				}
				score += n
			}
		}
	}
	currNum, err := strconv.Atoi(current)
	if err != nil {
		log.Fatal(err)
	}
	score *= currNum
	return score
}

func main() {
	// variables + data structures to track result value
	var numbers []string
	var boards []*Board
	var tmpBoards []*Board
	called := map[string]bool{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	numbers = strings.Split(sc.Text(), ",")
	boards = loadBoards(sc)

	for _, num := range numbers {
		called[num] = true
		if len(boards) == 1 {
			if boards[0].checkColumns(called) || boards[0].checkRows(called) {
				fmt.Printf("last board %v\n", boards[0])
				fmt.Printf("last board score %d\n", boards[0].calculateScore(called, num))
				return
			}
		}
		for _, b := range boards {
			if !b.checkColumns(called) && !b.checkRows(called) {
				tmpBoards = append(tmpBoards, b)
			}
		}
		boards = tmpBoards
		tmpBoards = nil
	}
}
