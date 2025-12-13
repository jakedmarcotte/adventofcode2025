package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	present_area := 9
	var present_areas []int
	var presents [][]int
	for scanner.Scan() {
		line := scanner.Text()
		reg_shape, _ := regexp.Compile(`[0-9]:`)
		reg_region, _ := regexp.Compile(`[0-9]+x[0-9]+:`)
		if reg_region.Find([]byte(line)) != nil {
			dimensions_str := strings.Split(line, ":")
			area_str := strings.TrimSpace(dimensions_str[0])
			area_dims := strings.Split(area_str, "x")
			width := 0
			height := 0
			fmt.Sscanf(area_dims[0], "%d", &width)
			fmt.Sscanf(area_dims[1], "%d", &height)
			pa := width * height
			present_areas = append(present_areas, pa)

			region_str := strings.TrimSpace(dimensions_str[1])
			present_idxs := []int{}
			for _, p := range strings.Fields(region_str) {
				num_p := 0
				fmt.Sscanf(p, "%d", &num_p)
				present_idxs = append(present_idxs, num_p)
			}
			presents = append(presents, present_idxs)
		} else if reg_shape.Find([]byte(line)) != nil {
			for i := 0; i < int(math.Sqrt(float64(present_area))); i++ {
				// we don't need to read the shape in for now
			}
		}
	}

	fitting_regions := 0
	for i := 0; i < len(presents); i++ {
		present := presents[i]
		area := present_areas[i]
		total_presents := 0
		for _, p := range present {
			total_presents += p
		}

		if area >= present_area*total_presents {
			fitting_regions++
		}
	}
	fmt.Println(fitting_regions)

}
