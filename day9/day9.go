package main

import (
	"bufio"
	"math"
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

	var coords [][2]int
	for scanner.Scan() {
		line := scanner.Text()
		coords_str := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords_str[0])
		y, _ := strconv.Atoi(coords_str[1])
		coords = append(coords, [2]int{x, y})
	}

	biggest_area := 0
	for i := 0; i < len(coords)-1; i++ {
		a := coords[i]
		for j := i + 1; j < len(coords); j++ {
			b := coords[j]
			dx := int(math.Abs(float64(b[0]-a[0]))) + 1
			dy := int(math.Abs(float64(b[1]-a[1]))) + 1
			area := dx * dy
			if area > biggest_area {
				biggest_area = area
			}
		}
	}

	println("Biggest area:", biggest_area)

}
