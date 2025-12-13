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
	var ingredient_ids [][2]int
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			break
		}
		ingredient_range_str := strings.Split(line, "-")
		num1, _ := strconv.Atoi(ingredient_range_str[0])
		num2, _ := strconv.Atoi(ingredient_range_str[1])
		ingredient_ids = append(ingredient_ids, [2]int{num1, num2})
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

	fmt.Println("Fresh ingredient IDs:", fresh_ingredients)
	println("Total fresh ingredients found:", len(fresh_ingredients))

}
