package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	// load the list of rolls
	var number_list [][]int
	var operators []string
	var lines int = 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "*") || strings.Contains(line, "+") {
			operators = strings.Fields(line)
			break
		}
		lines++
		numbers := strings.Fields(line)
		numSlice := make([]int, 0, len(numbers))
		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			numSlice = append(numSlice, num)
		}
		number_list = append(number_list, numSlice)
	}

	var results []int
	for opIndex, ops := range operators {
		var result int
		if ops == "*" {
			result = 1
		} else if ops == "+" {
			result = 0
		}
		for i := 0; i < lines; i++ {
			value := number_list[i][opIndex]
			if ops == "*" {
				result *= value
			} else if ops == "+" {
				result += value
			}
		}
		results = append(results, result)
	}
	fmt.Println("Results:", results)

	var total int
	for _, res := range results {
		total += res
	}
	fmt.Println("Total:", total)
}
