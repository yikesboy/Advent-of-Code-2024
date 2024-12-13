package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const cA = 3
const cB = 1

func main() {
	clawMachines, err := parseMachines()
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}
	res1, err := partOne(clawMachines)
	if err != nil {
		fmt.Printf("Error in part one: %v\n", err)
		return
	}
	fmt.Printf("minimum tokens needed for all winnable prizes: %d\n", res1)
}

type ClawMachine struct {
	id int
	A  [2][2]int
	b  [2]int
}

func partOne(clawMachines []ClawMachine) (int, error) {
	totalCost := 0

	for _, cm := range clawMachines {
		bestCost := math.MaxInt32
		found := false
		searchRange := 100

		for a := 0; a <= searchRange; a++ {
			for b := 0; b <= searchRange; b++ {
				x := cm.A[0][0]*a + cm.A[1][0]*b
				y := cm.A[0][1]*a + cm.A[1][1]*b

				if x == cm.b[0] && y == cm.b[1] {
					cost := cA*a + cB*b
					if cost < bestCost {
						bestCost = cost
						found = true
					}
				}
			}
		}

		if found {
			totalCost += bestCost
		}
	}

	return totalCost, nil
}

func parseMachines() ([]ClawMachine, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return []ClawMachine{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var clawMachines []ClawMachine
	id := 0

	for {
		var lines [3]string
		endReached := false

		for i := 0; i < 3; i++ {
			if !scanner.Scan() {
				endReached = true
				break
			}
			lines[i] = scanner.Text()
		}

		if endReached {
			break
		}

		if scanner.Scan() && scanner.Text() != "" {
			return []ClawMachine{}, errors.New("line was expected to be empty")
		}

		clawMachine := ClawMachine{id: id}
		for i, line := range lines {
			numbers, err := numberParser(line)
			if err != nil {
				return []ClawMachine{}, err
			}

			if i < 2 {
				clawMachine.A[i][0] = numbers[0]
				clawMachine.A[i][1] = numbers[1]
			} else {
				clawMachine.b[0] = numbers[0]
				clawMachine.b[1] = numbers[1]
			}
		}

		clawMachines = append(clawMachines, clawMachine)
		id++
	}

	return clawMachines, nil
}

func numberParser(line string) ([2]int, error) {
	var numbers [2]int
	var currentNumber strings.Builder
	inNumber := false
	index := 0

	for _, char := range line {
		if unicode.IsDigit(char) {
			inNumber = true
			currentNumber.WriteRune(char)
		} else if inNumber {
			inNumber = false
			if num, err := strconv.Atoi(currentNumber.String()); err != nil {
				return [2]int{}, err
			} else {
				numbers[index] = num
				index++
			}
			currentNumber.Reset()
		}
	}

	if inNumber {
		if num, err := strconv.Atoi(currentNumber.String()); err != nil {
			return [2]int{}, err
		} else {
			numbers[index] = num
		}
	}

	return numbers, nil
}
