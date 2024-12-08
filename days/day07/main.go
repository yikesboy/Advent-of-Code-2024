package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  res1, err := partOne()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in part one: %v\n", err)
		return
	}
  fmt.Printf("total calibration result: %d\n", res1)
}

func partOne() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

  result := 0
	scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    parsedLine := strings.Split(line, ": ")

    desiredOutcome, err := strconv.Atoi(parsedLine[0])
    if err != nil {
      return 0, err
    }

    equationNumberstrings := strings.Split(parsedLine[1], " ")

    equationNumbers := make([]int, 0)

    for _, op := range equationNumberstrings {
      if num, err := strconv.Atoi(op); err != nil {
        return 0, err
      } else {
        equationNumbers = append(equationNumbers, num)
      }
    }

    if isValid := checkValid(desiredOutcome, equationNumbers); isValid {
      result += desiredOutcome
    }
  }

  return result, nil
}

func checkValid(desiredOutcome int, equationNumbers []int) bool {
    operators := make([]string, len(equationNumbers)-1)
    numberCount := len(equationNumbers)
    /*
            0000001 -> 1
    << 1 -> 0000010 -> 2
    << 2 -> 0000100 -> 4
    << 3 -> 0001000 -> 8
    */
    maxCombinations := 1 << (numberCount - 1)

    for i := 0; i < maxCombinations; i++ {
        for pos := 0; pos < numberCount-1; pos++ {
            if (i & (1 << pos)) != 0 {
                operators[pos] = "*"
            } else {
                operators[pos] = "+"
            }
        }

        actualOutcome := equationNumbers[0]
        for j := 0; j < len(operators); j++ {
            if operators[j] == "+" {
                actualOutcome += equationNumbers[j+1]
            } else {
                actualOutcome *= equationNumbers[j+1]
            }
        }

        if actualOutcome == desiredOutcome {
            return true
        }
    }
    return false
}
