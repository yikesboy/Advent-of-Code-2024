package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readGrid()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading grid: %v\n", err)
		return
	}
	res1, err := partOne(grid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in part one: %v\n", err)
		return
	}
	fmt.Printf("Part 1 - XMAS count: %d\n", res1)

	res2, err := partTwo(grid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in part two: %v\n", err)
		return
	}
	fmt.Printf("Part 2 - X-MAS count: %d\n", res2)
}

func partOne(grid [][]byte) (int, error) {
	rows, cols := len(grid), len(grid[0])
	count := 0

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if isValidPos(i+3*dx, j+3*dy, rows, cols) {
					if grid[i][j] == 'X' &&
						grid[i+dx][j+dy] == 'M' &&
						grid[i+2*dx][j+2*dy] == 'A' &&
						grid[i+3*dx][j+3*dy] == 'S' {
						count++
					}
				}
			}
		}
	}

	return count, nil
}

func partTwo(grid [][]byte) (int, error) {
	rows, cols := len(grid), len(grid[0])
	count := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if isXMAS(grid, i, j, rows, cols) {
				count++
			}
		}
	}

	return count, nil
}

func readGrid() ([][]byte, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(grid) == 0 {
		return nil, fmt.Errorf("empty input file")
	}

	return grid, nil
}

func isValidPos(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func isXMAS(grid [][]byte, centerRow, centerCol, rows, cols int) bool {
	if centerRow == 0 || centerRow >= rows-1 || centerCol == 0 || centerCol >= cols-1 {
		return false
	}

	if isMAS(grid[centerRow-1][centerCol-1], grid[centerRow][centerCol], grid[centerRow+1][centerCol+1]) ||
		isMAS(grid[centerRow+1][centerCol+1], grid[centerRow][centerCol], grid[centerRow-1][centerCol-1]) {
		if isMAS(grid[centerRow-1][centerCol+1], grid[centerRow][centerCol], grid[centerRow+1][centerCol-1]) ||
			isMAS(grid[centerRow+1][centerCol-1], grid[centerRow][centerCol], grid[centerRow-1][centerCol+1]) {
			return true
		}
	}

	return false
}

func isMAS(first, middle, last byte) bool {
	return (first == 'M' && middle == 'A' && last == 'S') ||
		(first == 'S' && middle == 'A' && last == 'M')
}
