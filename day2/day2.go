package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hasRepeatingPattern(s string) bool {
	// iterate over first half of string
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		pattern := s[:i]
		repeated := true
		for j := i; j < len(s); j += i {
			if s[j:j+i] != pattern {
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}
	return false
}

func part1(invalid_id_sum *int, id_ranges []string) {
	for _, id_range := range id_ranges {
		var start, end int
		fmt.Sscanf(id_range, "%d-%d", &start, &end)
		for j := start; j <= end; j++ {
			id_str := strconv.Itoa(j)
			intLength := len(id_str)
			if intLength%2 == 0 {
				midpoint := intLength / 2
				firstHalf := id_str[:midpoint]
				secondHalf := id_str[midpoint:]
				if firstHalf == secondHalf {
					*invalid_id_sum += j
				}
			}

		}
	}
}

func part2(invalid_id_sum *int, id_ranges []string) {
	for _, id_range := range id_ranges {
		var start, end int
		fmt.Sscanf(id_range, "%d-%d", &start, &end)
		for j := start; j <= end; j++ {
			id_str := strconv.Itoa(j)
			if hasRepeatingPattern(id_str) {
				*invalid_id_sum += j
			}
		}
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		invalid_id_sum_p1 := 0
		invalid_id_sum_p2 := 0
		id_list := scanner.Text()
		id_ranges := strings.Split(id_list, ",")

		part1(&invalid_id_sum_p1, id_ranges)
		part2(&invalid_id_sum_p2, id_ranges)

		fmt.Println("Twice Repeating Invalid", invalid_id_sum_p1)
		fmt.Println("2+ Repeating Invalid", invalid_id_sum_p2)
	}

}
