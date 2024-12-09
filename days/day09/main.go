package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := partOne()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error analyzing memory file: %v\n", err)
		return
	}
	fmt.Printf("checksum: %d\n", num)
}

func partOne() (int, error) {
	memoryFile, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer memoryFile.Close()
	scanner := bufio.NewScanner(memoryFile)

	sum := 0
	if scanner.Scan() {
		input := scanner.Text()

		lengths := make([]int, len(input))
		for i, ch := range input {
			var err error
			lengths[i], err = strconv.Atoi(string(ch))
			if err != nil {
				return 0, err
			}
		}

		var blocks []int
		var freeSpaces []int
		fileID := 0

		for i := 0; i < len(lengths); i++ {
			length := lengths[i]

			if i%2 == 0 {
				for j := 0; j < length; j++ {
					blocks = append(blocks, fileID)
				}
				fileID++
			} else {
				startPos := len(blocks)
				for j := 0; j < length; j++ {
					blocks = append(blocks, -1)
					freeSpaces = append(freeSpaces, startPos+j)
				}
			}
		}

		result := make([]int, len(blocks))
		copy(result, blocks)

		freeSpaceIndex := 0
		for i := len(result) - 1; i >= 0 && freeSpaceIndex < len(freeSpaces); i-- {
			if result[i] != -1 && freeSpaces[freeSpaceIndex] < i {
				result[freeSpaces[freeSpaceIndex]] = result[i]
				result[i] = -1
				freeSpaceIndex++
			}
		}

		for pos, fileID := range result {
			if fileID != -1 {
				sum += pos * fileID
			}
		}
	}

	return sum, nil
}
