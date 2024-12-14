package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
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
}

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func partOne(robots []Robot) (int, error) {
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
			calculatePartition(robotCopies[th], start, end)
		}(th, start, end)
	}

	wg.Wait()

	finalRobots := make([]Robot, len(robots))
	for th := 0; th < threadCount; th++ {
		start := th * partitionSize
		end := start + partitionSize

		copy(finalRobots[start:end], robotCopies[th][start:end])
	}

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

func calculatePartition(robots []Robot, start int, end int) {
	for sec := 0; sec < 100; sec++ {
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
