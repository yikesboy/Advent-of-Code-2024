package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num1, err := partOne()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in partOne(): %v\n", err)
		return
	}
	fmt.Printf(": %d\n", num1)
}

func partOne() (int, error) {

	inputString := "1 24596 0 740994 60 803 8918 9405859"

	var stones []string = strings.Split(inputString, " ")

	for blink := 0; blink < 25; blink++ {

		var newStones []string
		for _, stone := range stones {
			length := len(stone)

			if stone == "0" {
				newStones = append(newStones, "1")
			} else if length%2 == 0 {
				mid := length / 2
				firstHalf := stone[:mid]
				secondHalf := stone[mid:]

				secondHalfInt, err := strconv.Atoi(secondHalf)
				if err != nil {
					return 0, err
				}

				newStones = append(newStones, firstHalf)
				newStones = append(newStones, strconv.Itoa(secondHalfInt))

			} else {
				stoneInt, err := strconv.Atoi(stone)
				if err != nil {
					return 0, err
				}

				stoneInt *= 2024
				stoneStr := strconv.Itoa(stoneInt)
				newStones = append(newStones, stoneStr)
			}
		}
		stones = newStones
	}

	return len(stones), nil
}