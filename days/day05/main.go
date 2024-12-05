package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	res1, err := partOne()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in part one: %v\n", err)
		return
	}
	fmt.Printf("sum of valid middle page numbers: %d\n", res1)
}

func partOne() (int, error) {

	file, err := os.Open("input.txt")
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	rules := make(map[string][]string)
	result := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parsedLine := strings.Split(line, "|")
		rules[parsedLine[0]] = append(rules[parsedLine[0]], parsedLine[1])
	}

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := strings.Split(line, ",")
		if checkPageOrdering(rules, parsedLine) {
			if middleNum, err := strconv.Atoi(parsedLine[len(parsedLine)/2]); err != nil {
				return 0, nil
			} else {
				result += middleNum
			}
		}

	}
	return result, nil
}

func checkPageOrdering(rules map[string][]string, pageOrder []string) bool {
	positions := make(map[string]int)

	for i, page := range pageOrder {
		positions[page] = i
	}

	for _, page := range pageOrder {
		if follows, ok := rules[page]; ok {
			for _, mustFollow := range follows {
				if pos, exists := positions[mustFollow]; exists {
					if positions[page] >= pos {
						return false
					}
				}
			}
		}
	}

	return true
}
