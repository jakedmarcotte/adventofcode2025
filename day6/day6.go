package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) {
	// load the list of rolls
	var number_str_list [][]string
	var max_lengths []int
	var operators []string
	for _, line := range lines {
		if strings.Contains(line, "*") || strings.Contains(line, "+") {
			operators = strings.Fields(line)
			break
		}
		numbers := strings.Fields(line)
		numSlice := make([]int, 0, len(numbers))
		max_length := 0
		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			numSlice = append(numSlice, num)
			if len(numStr) > max_length {
				max_length = len(numStr)
			}
		}
		max_lengths = append(max_lengths, max_length)
		number_str_list = append(number_str_list, numbers)
	}

	var results []int
	for opIndex, ops := range operators {
		var result int
		if ops == "*" {
			result = 1
		} else if ops == "+" {
			result = 0
		}
		for i := 0; i < len(lines)-1; i++ {

			value, _ := strconv.Atoi(number_str_list[i][opIndex])
			if ops == "*" {
				result *= value
			} else if ops == "+" {
				result += value
			}
		}
		results = append(results, result)
	}

	var total int
	for _, res := range results {
		total += res
	}
	fmt.Println("Part 1 Total:", total)
}

func part2(lines []string) {
	var operators []byte
	ops := strings.Fields(lines[len(lines)-1])
	num_of_problems := len(ops)
	number_list := make([][]int, num_of_problems)
	current_problem_idx := 0
	for i := len(lines[0]) - 1; i >= 0; i-- {
		curr_num := ""
		next_problem := false
		// go character by character and add to arrays
		for j := 0; j <= len(lines)-1; j++ {
			char := lines[j][i]
			if char == ' ' {
				continue
			} else if char == '*' || char == '+' {
				operators = append(operators, char)
				next_problem = true
			} else {
				// process digit
				curr_num += string(char)
			}
		}
		num, _ := strconv.Atoi(curr_num)
		number_list[current_problem_idx] = append(number_list[current_problem_idx], num)
		if next_problem {
			current_problem_idx++
			i--
			next_problem = false
		}
	}

	var results []int
	for i := 0; i < num_of_problems; i++ {
		var result int
		if operators[i] == '*' {
			result = 1
		} else if operators[i] == '+' {
			result = 0
		}
		for j := 0; j < len(number_list[i]); j++ {
			value := number_list[i][j]
			if operators[i] == '*' {
				result *= value
			} else if operators[i] == '+' {
				result += value
			}
		}
		results = append(results, result)
	}

	var total int
	for _, res := range results {
		total += res
	}
	fmt.Println("Part 2 Total:", total)
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	// load the list of rolls
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}

	part1(lines)
	part2(lines)
}
