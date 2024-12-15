package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wareHouseMap, moves, robotStart, err := parseInput()
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}

	num1, err := partOne(wareHouseMap, moves, robotStart)
	if err != nil {
		fmt.Printf("Error in part one: %v\n", err)
		return
	}
	fmt.Printf("sum of gps coordinates %d\n", num1)
}

func partOne(wareHouseMap [][]string, moves []string, robotStart [2]int) (int, error) {
	robotPos := robotStart
	directions := map[string][2]int{
		"<": {0, -1},
		"^": {-1, 0},
		"v": {1, 0},
		">": {0, 1},
	}

	for _, move := range moves {
		if dir, ok := directions[move]; ok {
			calcMove(wareHouseMap, dir, &robotPos)
		}
	}

	sum := 0
	boxCount := 0
	for y := 0; y < len(wareHouseMap); y++ {
		for x := 0; x < len(wareHouseMap[y]); x++ {
			if wareHouseMap[y][x] == "O" {
				boxCount++
				gps := (y * 100) + x
				sum += gps
			}
		}
	}

	return sum, nil
}

func calcMove(wareHouseMap [][]string, direction [2]int, robotPos *[2]int) {
	newY := robotPos[0] + direction[0]
	newX := robotPos[1] + direction[1]

	if newY < 0 || newY >= len(wareHouseMap) || newX < 0 || newX >= len(wareHouseMap[0]) {
		return
	}

	nextCell := wareHouseMap[newY][newX]

	if nextCell == "." {
		wareHouseMap[robotPos[0]][robotPos[1]] = "."
		wareHouseMap[newY][newX] = "@"
		robotPos[0] = newY
		robotPos[1] = newX
		return
	}

	if nextCell == "O" {
		boxes := [][2]int{{newY, newX}}
		checkY, checkX := newY+direction[0], newX+direction[1]

		for checkY >= 0 && checkY < len(wareHouseMap) &&
			checkX >= 0 && checkX < len(wareHouseMap[0]) {

			if wareHouseMap[checkY][checkX] == "O" {
				boxes = append(boxes, [2]int{checkY, checkX})
				checkY += direction[0]
				checkX += direction[1]
			} else {
				break
			}
		}

		if checkY < 0 || checkY >= len(wareHouseMap) ||
			checkX < 0 || checkX >= len(wareHouseMap[0]) ||
			wareHouseMap[checkY][checkX] == "#" {
			return
		}

		if wareHouseMap[checkY][checkX] == "." {
			wareHouseMap[checkY][checkX] = "O"
			for i := len(boxes) - 1; i > 0; i-- {
				wareHouseMap[boxes[i][0]][boxes[i][1]] = "O"
			}

			wareHouseMap[newY][newX] = "@"
			wareHouseMap[robotPos[0]][robotPos[1]] = "."
			robotPos[0] = newY
			robotPos[1] = newX
		}
	}
}

func parseInput() ([][]string, []string, [2]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return [][]string{}, []string{}, [2]int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var warehouseMap [][]string
	var moves []string
	var robotStart [2]int

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		row := strings.Split(line, "")
		for x, elem := range row {
			if elem == "@" {
				robotStart[0] = y
				robotStart[1] = x
			}
		}
		warehouseMap = append(warehouseMap, row)
		y++
	}

	for scanner.Scan() {
		line := scanner.Text()
		parsedMoves := strings.Split(line, "")

		moves = append(moves, parsedMoves...)
	}

	return warehouseMap, moves, robotStart, nil
}
