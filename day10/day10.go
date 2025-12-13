package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	var indicators [][]rune
	buttons := [][][]int{}
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)

		button_list := [][]int{}
		// iterate over each token
		for _, token := range tokens {
			if strings.Contains(token, "[") {
				indicator := strings.TrimFunc(token, func(r rune) bool {
					return r == '[' || r == ']'
				})
				indicators = append(indicators, []rune(indicator))
			} else if strings.Contains(token, "(") {
				button_str := strings.TrimFunc(token, func(r rune) bool {
					return r == '(' || r == ')'
				})
				button_idxs := []int{}
				for _, b := range strings.Split(button_str, ",") {
					b_val, _ := strconv.Atoi(b)
					button_idxs = append(button_idxs, b_val)
				}
				button_list = append(button_list, button_idxs)
			} else if strings.Contains(token, "{") {
				// ignore voltage for now
				buttons = append(buttons, button_list)
			}
		}
		index++
	}

	fewest_presses := 0
	for idx, indicator := range indicators {
		b := buttons[idx]
		initial_light := []rune{}
		for r := 0; r < len(indicator); r++ {
			initial_light = append(initial_light, '.')
		}
		min_presses := len(indicator)
		for i := 1; i < int(math.Pow(2, float64(len(b)))); i++ {
			var moves [][]int
			presses := 0
			light := make([]rune, len(indicator))
			copy(light, initial_light)

			binaryString := strconv.FormatInt(int64(i), 2)
			padding := strings.Repeat("0", len(b)-len(binaryString))

			for b_idx, bin := range padding + binaryString {
				if int(bin-'0') == 1 {
					btn := b[b_idx]
					moves = append(moves, btn)
					for _, s := range btn {
						curr := light[s]
						if curr == '.' {
							light[s] = '#'
						} else {
							light[s] = '.'
						}
					}
					presses++
				}
			}
			if slices.Equal(light, indicator) && presses < min_presses {
				min_presses = presses
			}
		}
		fewest_presses += min_presses
	}
	fmt.Println(fewest_presses)
}
