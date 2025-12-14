package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func visitor(row int, col int, rows [][]rune, visited *map[string]bool, accessible *[][2]int) {
	key := fmt.Sprintf("%d,%d", row, col)
	if _, exists := (*visited)[key]; exists {
		return
	}
	(*visited)[key] = true

	// visit neighbors (up, down, left, right)
	directions := [][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	diagonals := [][2]int{
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	if rows[row][col] == '@' {
		adjacents := 0

		// check adjacent cells
		for _, dir := range append(directions, diagonals...) {
			newRow := row + dir[0]
			newCol := col + dir[1]
			if newRow >= 0 && newRow < len(rows) && newCol >= 0 && newCol < len(rows[0]) && rows[newRow][newCol] == '@' {
				adjacents++
			}
		}

		if adjacents < 4 {
			*accessible = append(*accessible, [2]int{row, col})
		}
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newRow < len(rows) && newCol >= 0 && newCol < len(rows[0]) {
			visitor(newRow, newCol, rows, visited, accessible)
		}
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	// load the list of rolls
	var rows [][]rune
	for scanner.Scan() {
		rolls_str := scanner.Text()
		fmt.Println(rolls_str)
		var row []rune
		for _, char := range rolls_str {
			row = append(row, char)
		}

		rows = append(rows, row)
	}

	// remove the accessible positions from the grid
	no_removals := false
	total_accessible := 0
	for !no_removals {
		// inital rows state
		fmt.Println("Current grid state:")
		init_rows := make([][]rune, len(rows))
		for r, row := range rows {
			init_rows[r] = make([]rune, len(row))
			copy(init_rows[r], row)
			fmt.Println(string(init_rows[r]))
		}

		visited := make(map[string]bool)
		var accessible [][2]int
		visitor(0, 0, rows, &visited, &accessible)
		removing := len(accessible)
		total_accessible += removing
		fmt.Println("Removing", removing)

		// clearing accessible positions
		for _, pos := range accessible {
			rows[pos[0]][pos[1]] = '.'
		}

		fmt.Println("Next grid state:")
		for _, row := range rows {
			fmt.Println(string(row))
		}

		if reflect.DeepEqual(rows, init_rows) {
			no_removals = true
		}
	}
	fmt.Println("Total accessible positions after removals:", total_accessible)
}
