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
	fmt.Printf("safe report count: %d\n", res1)

	res2, err := partTwo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in part two: %v\n", err)
		return
	}
	fmt.Printf("safe report count with problem dampener: %d\n", res2)
}

func partOne() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		safe, err := isSafe(numbers)
		if err != nil {
			return 0, err
		}
		if safe {
			safeCount++
		}
	}
	return safeCount, nil
}

func partTwo() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())

		if len(numbers) < 2 {
			continue
		}

		safe, err := isSafe(numbers)
		if err != nil {
			return 0, err
		}
		if safe {
			safeCount++
			continue
		} else {
			for i := range numbers {
				withoutIndex := make([]string, 0, len(numbers)-1)
				withoutIndex = append(withoutIndex, numbers[:i]...)
				withoutIndex = append(withoutIndex, numbers[i+1:]...)

				safe, err := isSafe(withoutIndex)
				if err != nil {
					return 0, err
				}
				if safe {
					safeCount++
					break
				}
			}
		}
	}
	return safeCount, nil
}

func isSafe(numbers []string) (bool, error) {
	if len(numbers) < 2 {
		return false, nil
	}

	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		return false, err
	}

	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		return false, err
	}

	isIncreasing := num2 > num1

	for i := 1; i < len(numbers); i++ {
		newnum, err := strconv.Atoi(numbers[i])
		if err != nil {
			return false, err
		}

		diff := newnum - num1
		if isIncreasing {
			if diff <= 0 || diff > 3 {
				return false, nil
			}
		} else {
			if diff >= 0 || diff < -3 {
				return false, nil
			}
		}
		num1 = newnum
	}

	return true, nil
}
