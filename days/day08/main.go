package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    antinodeCount, err := countAntinodes()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error analyzing antenna map: %v\n", err)
        return
    }
    fmt.Printf("total unique antinode locations: %d\n", antinodeCount)
}

func countAntinodes() (int, error) {
    antennaMap, err := os.Open("input.txt")
    if err != nil {
        return 0, err
    }
    defer antennaMap.Close()

    scanner := bufio.NewScanner(antennaMap)
    frequencyLocations := make(map[string][]Position)
    mapBoundary := Position{0, 0}

    for y := 0; scanner.Scan(); y++ {
        line := scanner.Text()
        for x, c := range []byte(line) {
            if string(c) != "." {
                frequencyLocations[string(c)] = append(frequencyLocations[string(c)], Position{x, y})
            }
            mapBoundary.x = x
        }
        mapBoundary.y = y
    }
    if err := scanner.Err(); err != nil {
        return 0, err
    }

    antinodes := make(map[string]struct{})
    for frequency := range frequencyLocations {
        antennas := frequencyLocations[frequency]
        if err := findAntinodes(antennas, mapBoundary, antinodes); err != nil {
            return 0, err
        }
    }

    return len(antinodes), nil
}

func findAntinodes(antennas []Position, mapBoundary Position, antinodes map[string]struct{}) error {
    for i := 0; i < len(antennas); i++ {
        for j := i + 1; j < len(antennas); j++ {
            antinode1 := Position{2*antennas[i].x - antennas[j].x, 2*antennas[i].y - antennas[j].y}
            if antinode1.withinMapBounds(mapBoundary) {
                antinodes[antinode1.ToString()] = struct{}{}
            }
            antinode2 := Position{2*antennas[j].x - antennas[i].x, 2*antennas[j].y - antennas[i].y}
            if antinode2.withinMapBounds(mapBoundary) {
                antinodes[antinode2.ToString()] = struct{}{}
            }
        }
    }
    return nil
}

type Position struct {
    x, y int
}

func (p *Position) ToString() string {
    return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *Position) withinMapBounds(boundary Position) bool {
    return p.x >= 0 && p.x <= boundary.x && p.y >= 0 && p.y <= boundary.y
}
