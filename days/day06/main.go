package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"time"
)

func main() {
	res1, err := partOne()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in part one: %v\n", err)
		return
	}
	fmt.Printf("(distinct) visited fields: %d\n", res1)
	fmt.Println("-----------------------------------------------------")
}

func partOne() (int, error) {

	var visitedFields = make(map[string]bool)
	guardDirection := "^"
	puzzleMap, guardPos, err := parseGrid()
	if err != nil {
		return 0, err
	}

	x, y := guardPos[0], guardPos[1]
	for {

		if y >= len(puzzleMap) || y <= 0 || x >= len(puzzleMap[0]) || x <= 0 {
			break
		}

		//time.Sleep(100 * time.Millisecond)
		printMapState(puzzleMap)

		switch guardDirection {
		case "^":
			for y-1 > -1 && puzzleMap[y-1][x] != "#" {
				puzzleMap[y][x] = "X"
				markVisited(&visitedFields, x, y)
				y--
			}
			guardDirection = ">"

		case "<":
			for x-1 > -1 && puzzleMap[y][x-1] != "#" {
				puzzleMap[y][x] = "X"
				markVisited(&visitedFields, x, y)
				x--
			}
			guardDirection = "^"

		case ">":
			for x+1 < len(puzzleMap[y]) && puzzleMap[y][x+1] != "#" {
				puzzleMap[y][x] = "X"
				markVisited(&visitedFields, x, y)
				x++
			}
			guardDirection = "v"

		case "v":
			for y+1 < len(puzzleMap) && puzzleMap[y+1][x] != "#" {
				puzzleMap[y][x] = "X"
				markVisited(&visitedFields, x, y)
				y++
			}
			guardDirection = "<"
		}
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Guard End-Position: %d - %d\n", x, y)
	return len(visitedFields), nil
}

func parseGrid() ([][]string, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	puzzleMap := make([][]string, 0)
	lineNr := 0
	guardPos := []int{0, 0}

	for scanner.Scan() {
		lineNr++
		line := scanner.Text()
		parsedLine := strings.Split(line, "")
		for pos, item := range parsedLine {
			if item == "^" {
				guardPos[0] = pos
				guardPos[1] = lineNr
			}
		}
		puzzleMap = append(puzzleMap, parsedLine)
	}

	fmt.Printf("Guard Start-Position: %d - %d\n", guardPos[0], guardPos[1])

	return puzzleMap, guardPos, nil
}

func markVisited(visitedFields *map[string]bool, x int, y int) {
	key := fmt.Sprintf("%d,%d", x, y)
	(*visitedFields)[key] = true
}

func printMapState(puzzleMap [][]string) {
	fmt.Print("\033[H\033[2J")
	for y := 0; y < len(puzzleMap); y++ {
		for x := 0; x < len(puzzleMap[y]); x++ {
			fmt.Print(puzzleMap[y][x])
		}
		fmt.Println()
	}
}
