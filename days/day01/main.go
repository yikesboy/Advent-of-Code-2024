package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
	}
	defer file.Close()

	var list1 = make([]int, 0, 1000)
	var list2 = make([]int, 0, 1000)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 2 {
			continue
		}

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			continue
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	// Part 1
	distance := 0

	for i := range list1 {
		num1 := list1[i]
		num2 := list2[i]
		if num1 >= num2 {
			distance += num1 - num2
		} else {
			distance += num2 - num1
		}
	}

	fmt.Printf("distance between the historian's lists: %d\n", distance)

	// Part 2
	countMap1 := make(map[int]int)
	countMap2 := make(map[int]int)

	for i := range list1 {
		countMap1[list1[i]]++
		countMap2[list2[i]]++
	}

	similarityScore := 0
	for num, _ := range countMap1 {
		similarityScore += num * countMap2[num]
	}

	fmt.Printf("similarity score: %d\n", similarityScore)
}