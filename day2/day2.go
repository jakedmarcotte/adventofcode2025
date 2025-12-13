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

	for scanner.Scan() {
		invalid_id_sum := 0
		id_list := scanner.Text()
		id_ranges := strings.Split(id_list, ",")

		for i, id_range := range id_ranges {
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
						invalid_id_sum += j
					}
				}

			}
			fmt.Println(id_ranges[i])
		}
		fmt.Println(invalid_id_sum)
	}

}
