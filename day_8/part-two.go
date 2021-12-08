package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func main() {
	var sum int
	// Make an array that maps the indices to the segments in the index digit
	segments := [10]string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}

	// Make some maps that will be useful for lookup later
	segStringMap := make(map[string]string)
	for i, s := range segments {
		segStringMap[s] = strconv.Itoa(i)
	}
	segCountMap := make(map[rune]int)
	for _, s := range segments {
		for _, c := range s {
			segCountMap[c] += 1
		}
	}
	countSegMap := make(map[int][]rune)
	for k, v := range segCountMap {
		countSegMap[v] = append(countSegMap[v], k)
	}

	// Open and loop through lines of file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		// parse line
		row := strings.Split(sc.Text(), " | ")
		outputs := strings.Fields(row[1])
		inputs := strings.Fields(row[0])

		// More maps yay
		segMap := make(map[rune][]rune)
		segCountRowMap := make(map[rune]int)

		// These will be useful later
		var one string
		var four string
		var seven string

		for _, s := range inputs {
			switch len(s) {
			case len(segments[1]):
				one = s
			case len(segments[4]):
				four = s
			case len(segments[7]):
				seven = s
			}
			for _, c := range s {
				segCountRowMap[c] += 1
			}
		}
		// generate preliminary map by comparing the number of digits each segment appears in
		// we'll do some further deductions later
		for k, v := range segCountRowMap {
			segMap[k] = countSegMap[v]
		}

		// if there are multiple candidates for any segment, determine the correct mapping using known digits
		for k, v := range segMap {
			if len(v) == 2 {
				if contains(v, 'd') {
					if strings.ContainsRune(four, k) {
						segMap[k] = []rune{'d'}
					} else {
						segMap[k] = []rune{'g'}
					}
				} else if contains(v, 'a') {
					if strings.ContainsRune(seven, k) && !strings.ContainsRune(one, k) {
						segMap[k] = []rune{'a'}
					} else {
						segMap[k] = []rune{'c'}
					}
				}
			}
		}

		// decode the output string and add it to the sum
		var output string
		for _, s := range outputs {
			var digit string
			for _, c := range s {
				digit += string(segMap[c][0])
			}
			output += segStringMap[sortString(digit)]
		}

		outputInt, err := strconv.Atoi(output)
		if err != nil {
			log.Fatal(err)
		}
		sum += outputInt

	}

	fmt.Printf("result: %d \n", sum)
}
