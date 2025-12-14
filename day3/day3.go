package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// TODO: NOTE
// the key here is to go from index 0 to 3, determined by length of the battery - digits needed, rolling window function
func findIndexOfHighest(battery_bank []rune, start_index int, end_index int) int {
	highest := 0
	indexOfHighest := 0
	for i := start_index; i <= end_index; i++ {
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
		batteries = append(batteries, []rune(battery_str))
	}

	var voltage int
	var voltages []int
	for _, bank := range batteries {
		charge_indexes := []int{}

		// changing this to be 2 or 12 depending on part 1 or part 2
		digits := 12 // or 2
		curr := 0
		for i := 0; i < digits; i++ {
			index := findIndexOfHighest(bank, curr, len(bank)-digits+i)
			charge_indexes = append(charge_indexes, index)
			curr = index + 1
		}

		var voltage_str string
		for _, ci := range charge_indexes {
			voltage_str += string(bank[ci])
		}
		voltage, _ := strconv.Atoi(voltage_str)
		voltages = append(voltages, voltage)
	}

	for _, charge := range voltages {
		voltage += charge
	}
	fmt.Println("Total Voltage:", voltage)
}
