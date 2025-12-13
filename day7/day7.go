package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	var hub [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		hub = append(hub, row)
	}

	var splits [][2]int
	for r, row := range hub {
		for c, char := range row {
			if char == 'S' {
				println("Found start")
				hub[r+1][c] = '|'
				continue
			} else if char == '|' && r+1 < len(hub) && c < len(hub) {
				if hub[r+1][c] == '^' {
					hub[r+1][c-1] = '|'
					hub[r+1][c+1] = '|'
					splits = append(splits, [2]int{r + 1, c})
				} else {
					hub[r+1][c] = '|'
				}
				continue
			}
		}

		// print the row
		for _, char := range hub[r] {
			print(string(char))
		}
		println("")
	}
	fmt.Println("Splits found at:", splits)
	println("Total splits found:", len(splits))
}
