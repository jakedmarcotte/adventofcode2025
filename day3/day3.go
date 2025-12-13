package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findIndexOfHighest(battery_bank []rune, start_index int, consider_last bool) int {
	highest := 0
	indexOfHighest := start_index
	var tail_offset int = 0
	if consider_last == false {
		tail_offset = 1
	}
	for i := start_index; i < len(battery_bank)-tail_offset; i++ {
		battery_v := int(battery_bank[i])
		if battery_v > highest {
			highest = battery_v
			indexOfHighest = i
		}
	}
	return indexOfHighest
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	// load the list of batteries
	var batteries [][]rune
	for scanner.Scan() {
		battery_str := scanner.Text()
		fmt.Println(battery_str)

		batteries = append(batteries, []rune(battery_str))
	}

	fmt.Println(batteries)

	var charges []int
	var voltage int
	for _, bank := range batteries {
		var highest int
		var indexOfHighest = findIndexOfHighest(bank, 0, false)
		nextHighestIndex := findIndexOfHighest(bank, indexOfHighest+1, true)
		battery_seq := []rune{bank[indexOfHighest], bank[nextHighestIndex]}
		converted, err := strconv.Atoi(string(battery_seq))
		if err != nil {
			panic(err)
		}
		highest = converted
		charges = append(charges, highest)
	}

	for _, charge := range charges {
		println(charge)
		voltage += charge
	}
	println("Total Voltage:", voltage)
}
