package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"sync"
	"time"
)

const spaceY = 103
const spaceX = 101
const threadCount = 5

func main() {
	robots, err := parseInput()
	if err != nil {
		fmt.Printf("Error parsing input: %v", err)
	}
	if num1, err := partOne(robots); err != nil {
		fmt.Printf("Error in part one: %v", err)
	} else {
		fmt.Printf("saftey factor after 100 seconds have elapsed: %d\n", num1)
	}
	if num2, err := partTwo(robots); err != nil {
		fmt.Printf("Error in part two: %v", err)
	} else {
		fmt.Printf("easter egg is at second: %d\n", num2)
	}
}

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func partOne(robots []Robot) (int, error) {

	finalRobots := travelToFuture(robots, 100)

	var quarters [4]int
	for _, rb := range finalRobots {
		if rb.x == 50 || rb.y == 51 {
			continue
		}

		if rb.x < 50 && rb.y < 51 {
			quarters[0]++
		} else if rb.x > 50 && rb.y < 51 {
			quarters[1]++
		} else if rb.x < 50 && rb.y > 51 {
			quarters[2]++
		} else if rb.x > 50 && rb.y > 51 {
			quarters[3]++
		}
	}

	return quarters[0] * quarters[1] * quarters[2] * quarters[3], nil
}

func partTwo(robots []Robot) (int, error) {
	type ClusterState struct {
		iteration int
		value     float64
		robots    []Robot
	}

	var states []ClusterState
	finalRobots := travelToFuture(robots, 3700)

	for i := 4000; i < 7500; i++ {
		fmt.Printf("Processing iteration %d\n", i)
		robotsCopy := make([]Robot, len(finalRobots))
		copy(robotsCopy, finalRobots)

		clusterValue := clusterValueForIter(finalRobots)
		states = append(states, ClusterState{
			iteration: i,
			value:     clusterValue,
			robots:    robotsCopy,
		})

		finalRobots = travelToFuture(finalRobots, 1)
	}

	sort.Slice(states, func(i, j int) bool {
		return states[i].value < states[j].value
	})

	for i := 0; i < 20 && i < len(states); i++ {
		printField(states[i].robots)
		fmt.Printf("\nIteration: %d, Cluster Value: %f\n", states[i].iteration, states[i].value)
		time.Sleep(2 * time.Second)
	}

	return states[0].iteration, nil
}

func clusterValueForIter(robots []Robot) float64 {
	distances := make([]float64, len(robots))

	for i, rb1 := range robots {
		minDist := math.MaxFloat64
		for j, rb2 := range robots {
			if i == j {
				continue
			}

			dist := math.Sqrt(
				math.Pow(float64(rb1.x-rb2.x), 2) + math.Pow(float64(rb1.y-rb2.y), 2),
			)

			if dist < minDist {
				minDist = dist
			}
		}
		distances[i] = minDist
	}

	sum := 0.0
	for _, d := range distances {
		sum += d
	}

	return sum / float64(len(distances))
}

func calculatePartition(robots []Robot, start int, end int, seconds int) {
	for sec := 0; sec < seconds; sec++ {
		for rnr := start; rnr < end; rnr++ {
			dx := robots[rnr].x + robots[rnr].vx
			dy := robots[rnr].y + robots[rnr].vy
			if dx >= spaceX {
				dx -= spaceX
			} else if dx < 0 {
				dx += spaceX
			}
			if dy >= spaceY {
				dy -= spaceY
			} else if dy < 0 {
				dy += spaceY
			}
			robots[rnr].x = dx
			robots[rnr].y = dy
		}
	}
}

func parseInput() ([]Robot, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return []Robot{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var robots []Robot

	regex := regexp.MustCompile(`-?\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		numStrings := regex.FindAllString(line, -1)
		if len(numStrings) != 4 {
			return []Robot{}, errors.New("expected 4 integers")
		}

		var nums []int
		for _, ns := range numStrings {
			if num, err := strconv.Atoi(ns); err != nil {
				return []Robot{}, err
			} else {
				nums = append(nums, num)
			}
		}

		robot := Robot{
			x:  nums[0],
			y:  nums[1],
			vx: nums[2],
			vy: nums[3],
		}

		robots = append(robots, robot)
	}

	return robots, nil
}

func travelToFuture(robots []Robot, seconds int) []Robot {
	robotCopies := make([][]Robot, threadCount)
	partitionSize := len(robots) / threadCount

	for th := 0; th < threadCount; th++ {
		robotCopies[th] = make([]Robot, len(robots))
		copy(robotCopies[th], robots)
	}

	var wg sync.WaitGroup
	wg.Add(threadCount)

	for th := 0; th < threadCount; th++ {
		start := th * partitionSize
		end := start + partitionSize
		if th == threadCount-1 {
			end = len(robots)
		}

		go func(th, start, end int) {
			defer wg.Done()
			calculatePartition(robotCopies[th], start, end, seconds)
		}(th, start, end)
	}

	wg.Wait()

	finalRobots := make([]Robot, len(robots))
	for th := 0; th < threadCount; th++ {
		start := th * partitionSize
		end := start + partitionSize

		copy(finalRobots[start:end], robotCopies[th][start:end])
	}

	return finalRobots
}

func printField(robots []Robot) {
	field := make([][]string, spaceY)
	for i := range field {
		field[i] = make([]string, spaceX)
		for j := range field[i] {
			field[i][j] = "."
		}
	}

	for _, rb := range robots {
		field[rb.y][rb.x] = "#"
	}

	fmt.Println("\033[2J")
	fmt.Println("\033[H")
	for _, row := range field {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}
