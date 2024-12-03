package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum1, err := partOne()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in part one: %v\n", err)
		return
	}
	fmt.Printf("valid mul operation sum %d\n", sum1)

	sum2, err := partTwo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in part two: %v\n", err)
		return
	}
	fmt.Printf("valid mul operation sum %d\n", sum2)
}

func partOne() (int, error) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		return 0, err
	}
	contentString := string(content)
	contentArray := strings.Split(contentString, "")

	result := 0
	i := 0
	for i < len(contentArray) {
		if i+3 >= len(contentArray) {
			break
		}

		if contentArray[i] == "m" && contentArray[i+1] == "u" && contentArray[i+2] == "l" && contentArray[i+3] == "(" {
			num1String := ""
			num2String := ""
			isValid := true
			subIndex := i + 4

			for subIndex < len(contentArray) {
				if contentArray[subIndex] == "," {
					subIndex++
					break
				}
				if _, err := strconv.Atoi(contentArray[subIndex]); err != nil {
					isValid = false
					break
				}
				num1String += contentArray[subIndex]
				subIndex++
			}

			if isValid && subIndex < len(contentArray) {
				for subIndex < len(contentArray) {
					if contentArray[subIndex] == ")" {
						subIndex++
						break
					}
					if _, err := strconv.Atoi(contentArray[subIndex]); err != nil {
						isValid = false
						break
					}
					num2String += contentArray[subIndex]
					subIndex++
				}
			}

			if isValid && len(num1String) > 0 && len(num1String) <= 3 && len(num2String) > 0 && len(num2String) <= 3 {
				num1, err1 := strconv.Atoi(num1String)
				num2, err2 := strconv.Atoi(num2String)
				if err1 == nil && err2 == nil {
					result += num1 * num2
				}
			}
			i = subIndex
		} else {
			i++
		}
	}
	return result, nil
}

func partTwo() (int, error) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		return 0, err
	}
	contentString := string(content)
	contentArray := strings.Split(contentString, "")

	result := 0
	isDo := true

	i := 0
	for i < len(contentArray) {
		if i+3 < len(contentArray) && contentArray[i] == "d" && contentArray[i+1] == "o" && contentArray[i+2] == "(" && contentArray[i+3] == ")" {
            isDo = true
            i += 4
            continue
        }

		if i+4 < len(contentArray) && contentArray[i] == "d" && contentArray[i+1] == "o" && contentArray[i+2] == "n" && contentArray[i+3] == "'" && contentArray[i+4] == "t" { 
			isDo = false
			i += 4
			continue
		}

		if i+3 >= len(contentArray) {
			break
		}

		if contentArray[i] == "m" && contentArray[i+1] == "u" && contentArray[i+2] == "l" && contentArray[i+3] == "(" {
			num1String := ""
			num2String := ""
			isValid := true
			subIndex := i + 4

			for subIndex < len(contentArray) {
				if contentArray[subIndex] == "," {
					subIndex++
					break
				}
				if _, err := strconv.Atoi(contentArray[subIndex]); err != nil {
					isValid = false
					break
				}
				num1String += contentArray[subIndex]
				subIndex++
			}

			if isValid && subIndex < len(contentArray) {
				for subIndex < len(contentArray) {
					if contentArray[subIndex] == ")" {
						subIndex++
						break
					}
					if _, err := strconv.Atoi(contentArray[subIndex]); err != nil {
						isValid = false
						break
					}
					num2String += contentArray[subIndex]
					subIndex++
				}
			}

			if isDo && isValid && len(num1String) > 0 && len(num1String) <= 3 && len(num2String) > 0 && len(num2String) <= 3 {
				num1, err1 := strconv.Atoi(num1String)
				num2, err2 := strconv.Atoi(num2String)
				if err1 == nil && err2 == nil {
					result += num1 * num2
				}
			}
			i = subIndex
		} else {
			i++
		}
	}
	return result, nil
}