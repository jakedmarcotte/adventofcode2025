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
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	// load the list of rolls
	var ingredient_ids [][2]int
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			break
		}
		ingredient_range_str := strings.Split(line, "-")
		num1, _ := strconv.Atoi(ingredient_range_str[0])
		num2, _ := strconv.Atoi(ingredient_range_str[1])
		if num1 < num2 {
			ingredient_ids = append(ingredient_ids, [2]int{num1, num2})
		} else {
			ingredient_ids = append(ingredient_ids, [2]int{num2, num1})
		}
	}

	var fresh_ingredients []int
	for scanner.Scan() {
		ingredient_id, _ := strconv.Atoi(scanner.Text())
		for _, id_range := range ingredient_ids {
			if ingredient_id >= id_range[0] && ingredient_id <= id_range[1] {
				println("Ingredient ID", ingredient_id, "is valid for range", id_range[0], "-", id_range[1])
				fresh_ingredients = append(fresh_ingredients, ingredient_id)
				break
			}
		}
	}

	// part 2 here
	// Sort ingredient_ids by the first element of each range
	sort.Slice(ingredient_ids, func(i, j int) bool {
		return ingredient_ids[i][0] < ingredient_ids[j][0]
	})

	low_bound := ingredient_ids[0][0]
	high_bound := ingredient_ids[0][1]
	possible_ids := ingredient_ids[0][1] - ingredient_ids[0][0] + 1
	for i := 0; i < len(ingredient_ids); i++ {
		min := ingredient_ids[i][0]
		max := ingredient_ids[i][1]

		if min >= low_bound && max <= high_bound {
			continue
		} else {

			if max > high_bound && min <= low_bound {
				possible_ids += max - high_bound
				high_bound = max
			} else if max > high_bound && min > low_bound && min > high_bound {
				possible_ids += max - min + 1
				high_bound = max
				low_bound = min
			} else if min > low_bound && max > high_bound {
				possible_ids += max - high_bound
				high_bound = max
				low_bound = min
			}
		}

		fmt.Println("Current range:", ingredient_ids[i], "Low bound:", low_bound, "High bound:", high_bound, "Possible IDs so far:", possible_ids)

	}

	println("Total fresh ingredients found:", len(fresh_ingredients))
	fmt.Println("Possible ingredient IDs:", possible_ids)

}
