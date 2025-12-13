package main

import (
	"bufio"
	"fmt"
	"os"
)

func visitor(row int, col int, rows [][]rune, visited *map[string]bool, accessible *[]string) {
	key := fmt.Sprintf("%d,%d", row, col)
	if _, exists := (*visited)[key]; exists {
		return
	}
	(*visited)[key] = true

	// process current position
	fmt.Printf("Visiting row %d, col %d: %c\n", row, col, rows[row][col])

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
			*accessible = append(*accessible, key)
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

	visited := make(map[string]bool)
	var accessible []string
	visitor(0, 0, rows, &visited, &accessible)

	fmt.Println(accessible)
	fmt.Println("Total accessible positions:", len(accessible))
}
