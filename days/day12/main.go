package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	garden, err := parseGarden()
	if err != nil {
		return
	}
	res1, err := partOne(garden)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in part one: %v\n", err)
		return
	}
	fmt.Printf("part one: %d\n", res1)
}

type gardenTile struct {
	id      int
	value   string
	visited bool
	x       int
	y       int
}

type RegionStat struct {
	area      int
	perimeter int
}

func partOne(garden [][]gardenTile) (int, error) {
	regionStats := make(map[int]*RegionStat)
	regionId := 0

	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[y]); x++ {
			if !garden[y][x].visited {
				regionId++
				stats := &RegionStat{}
				regionStats[regionId] = stats
				dfs(garden, x, y, regionId, garden[y][x].value, stats)
			}
		}
	}

	result := 0
	for _, stats := range regionStats {
		result += stats.area * stats.perimeter
	}
	return result, nil
}

func dfs(garden [][]gardenTile, x, y, regionId int, currentPlant string, stats *RegionStat) {
	if y < 0 || y >= len(garden) || x < 0 || x >= len(garden[y]) {
		return
	}

	tile := &garden[y][x]
	if tile.visited || tile.value != currentPlant {
		return
	}

	tile.id = regionId
	tile.visited = true
	stats.area++
	stats.perimeter += getPerimeterCount(garden, *tile)

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		dfs(garden, x+dir[0], y+dir[1], regionId, currentPlant, stats)
	}
}

func getPerimeterCount(garden [][]gardenTile, tile gardenTile) int {

	result := 0
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		dx := tile.x + dir[0]
		dy := tile.y + dir[1]

		if dx < 0 || dy < 0 || dy >= len(garden) || dx >= len(garden[0]) {
			result++
			continue
		}

		if garden[dy][dx].value != tile.value {
			result += 1
		}
	}

	return result
}

func parseGarden() ([][]gardenTile, error) {
	gardenFile, err := os.Open("input.txt")
	if err != nil {
		return [][]gardenTile{}, err
	}
	defer gardenFile.Close()

	scanner := bufio.NewScanner(gardenFile)
	garden := make([][]gardenTile, 0)

	y := 0
	for scanner.Scan() {
		gardenRow := scanner.Text()
		parsedGardenRow := strings.Split(gardenRow, "")
		row := make([]gardenTile, 0, len(parsedGardenRow))
		for x, plot := range parsedGardenRow {
			row = append(row, gardenTile{id: 0, value: plot, x: x, y: y})
		}
		garden = append(garden, row)
		y++
	}

	fmt.Println(len(garden))
	return garden, nil
}
