package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	hikingMap, trailStarts, err := parseMap()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input file: %v\n", err)
		return
	}
	if num, err := partOne(hikingMap, trailStarts); err != nil {
		fmt.Fprintf(os.Stderr, "error part one: %v\n", err)
		return
	} else {
		fmt.Printf("solution part one: %d\n", num)
	}
	if num, err := partTwo(hikingMap, trailStarts); err != nil {
		fmt.Fprintf(os.Stderr, "error part two: %v\n", err)
		return
	} else {
		fmt.Printf("solution part two: %d\n", num)
	}
}

func partOne(hikingMap [][]int, trailStarts []mapTile) (int, error) {
	result := 0

	for _, start := range trailStarts {
		reachableNines := make(map[string]bool)

		visited := make(map[string]bool)
		dfs(hikingMap, start, visited, &reachableNines)

		result += len(reachableNines)
	}

	return result, nil
}

func partTwo(hikingMap [][]int, trailStarts []mapTile) (int, error) {
	result := 0

	for _, start := range trailStarts {
		paths := countPaths(hikingMap, start)
		result += paths
	}

	return result, nil
}

func countPaths(hikinMap [][]int, start mapTile) int {
	cache := make(map[string]int)
	return dfsCountPaths(hikinMap, start, cache)
}

func dfsCountPaths(hikingMap [][]int, current mapTile, cache map[string]int) int {
	key := fmt.Sprintf("%d,%d", current.x, current.y)

	if count, exists := cache[key]; exists {
		return count
	}

	if current.value == 9 {
		return 1
	}

	totalPaths := 0

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		newX := current.x + dir[0]
		newY := current.y + dir[1]

		if newY >= 0 && newY < len(hikingMap) && newX >= 0 && newX < len(hikingMap[0]) {
			if hikingMap[newY][newX] == current.value+1 {
				totalPaths += dfsCountPaths(hikingMap, mapTile{newX, newY, hikingMap[newY][newX]}, cache)
			}
		}
	}

	cache[key] = totalPaths
	return totalPaths
}

func dfs(hikingMap [][]int, current mapTile, visited map[string]bool, reachableNines *map[string]bool) {
	key := fmt.Sprintf("%d,%d", current.x, current.y)

	if visited[key] {
		return
	}

	visited[key] = true

	if current.value == 9 {
		(*reachableNines)[key] = true
		return
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		newX := current.x + dir[0]
		newY := current.y + dir[1]

		if newX >= 0 && newX < len(hikingMap) && newY >= 0 && newY < len(hikingMap) {
			if hikingMap[newY][newX] == current.value+1 {
				dfs(hikingMap, mapTile{newX, newY, hikingMap[newY][newX]}, visited, reachableNines)
			}
		}
	}

}

func parseMap() ([][]int, []mapTile, error) {
	mapFile, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer mapFile.Close()

	var hikingMap [][]int
	var trailStarts []mapTile

	scanner := bufio.NewScanner(mapFile)
	y := 0
	for scanner.Scan() {
		input := scanner.Text()
		parsedLine := strings.Split(input, "")
		row := make([]int, len(parsedLine))

		for x, elem := range parsedLine {
			if num, err := strconv.Atoi(elem); err != nil {
				return nil, nil, err
			} else {
				row[x] = num
				if num == 0 {
					trailStarts = append(trailStarts, mapTile{x: x, y: y, value: 0})
				}
			}
		}
		hikingMap = append(hikingMap, row)
		y++
	}

	return hikingMap, trailStarts, nil
}

type mapTile struct {
	x     int
	y     int
	value int
}
