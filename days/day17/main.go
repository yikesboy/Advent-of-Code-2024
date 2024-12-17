package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	register, program, err := parseRegAndIns()
	if err != nil {
		fmt.Printf("error parsing input: %v\n", err)
		return
	}
	str, err := partOne(register, program)
	if err != nil {
		fmt.Printf("error in part one: %v\n", err)
		return
	}
	fmt.Println(str)
}

func partOne(register [3]uint64, program []int) (string, error) {
	IP := 0
	var output string
	var reg = register

	for IP < len(program)-1 {
		opcode := program[IP]
		OP := program[IP+1]

		comboOP := getComboValue(&reg, OP)

		switch opcode {
		case 0:
			adv(&reg, comboOP)
		case 1:
			bxl(&reg, OP)
		case 2:
			bst(&reg, comboOP)
		case 3:
			jnz(&reg, OP, &IP)
		case 4:
			bxc(&reg)
		case 5:
			out(comboOP, &output)
		case 6:
			bdv(&reg, comboOP)
		case 7:
			cdv(&reg, comboOP)
		}

		if opcode != 3 {
			IP += 2
		}
	}

	return output, nil
}

func getComboValue(register *[3]uint64, op int) uint64 {
	switch op {
	case 4:
		return register[0]
	case 5:
		return register[1]
	case 6:
		return register[2]
	default:
		return uint64(op)
	}
}

func adv(register *[3]uint64, comboOP uint64) {
	(*register)[0] >>= comboOP
}

func bxl(register *[3]uint64, OP int) {
	(*register)[1] ^= uint64(OP)
}

func bst(register *[3]uint64, comboOP uint64) {
	(*register)[1] = comboOP & 7
}

func jnz(register *[3]uint64, OP int, IP *int) {
	if (*register)[0] != 0 {
		*IP = OP
	} else {
		*IP += 2
	}
}

func bxc(register *[3]uint64) {
	(*register)[1] ^= (*register)[2]
}

func out(comboOP uint64, output *string) {
	if len(*output) > 0 {
		*output += ","
	}
	*output += strconv.Itoa(int(comboOP & 7))
}

func bdv(register *[3]uint64, comboOP uint64) {
	(*register)[1] = (*register)[0] >> comboOP
}

func cdv(register *[3]uint64, comboOP uint64) {
	(*register)[2] = (*register)[0] >> comboOP
}

func parseRegAndIns() ([3]uint64, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return [3]uint64{}, []int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var register [3]uint64
	var program []int

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parsedLine := strings.Split(line, ": ")

		num, err := strconv.ParseUint(parsedLine[1], 10, 64)
		if err != nil {
			return [3]uint64{}, []int{}, err
		}
		register[i] = num
		i++
	}

	if !scanner.Scan() {
		return [3]uint64{}, []int{}, fmt.Errorf("missing program line")
	}
	line := scanner.Text()
	parsedLine := strings.Split(line, ": ")
	numbers := strings.Split(parsedLine[1], ",")
	for _, str := range numbers {
		if num, err := strconv.Atoi(str); err != nil {
			return [3]uint64{}, []int{}, err
		} else {
			program = append(program, num)
		}
	}

	return register, program, nil
}
