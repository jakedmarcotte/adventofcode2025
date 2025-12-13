package main

import (
	adventheap "adventofcode2025/heap"
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	x float64
	y float64
	z float64
}

type Junction struct {
	a Box
	b Box
}

type Circuit struct {
	boxes []Box
}

func distance(a, b Box) float64 {
	x := math.Pow(float64(a.x-b.x), 2)
	y := math.Pow(float64(a.y-b.y), 2)
	z := math.Pow(float64(a.z-b.z), 2)

	return math.Sqrt(x + y + z)
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	var boxes []Box
	connections, _ := strconv.Atoi(os.Args[2])
	for scanner.Scan() {
		var line = scanner.Text()
		coordinates_str := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates_str[0])
		y, _ := strconv.Atoi(coordinates_str[1])
		z, _ := strconv.Atoi(coordinates_str[2])
		box := Box{x: float64(x), y: float64(y), z: float64(z)}
		boxes = append(boxes, box)
	}

	m := adventheap.New(func(a, b Junction) bool {
		return distance(a.a, a.b) < distance(b.a, b.b)
	})

	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			heap.Push(m, Junction{boxes[i], boxes[j]})
		}
	}

	circuits := make(map[Box]*Circuit)

	for i := 0; i < connections; i++ {
		if m.Len() == 0 {
			break
		}

		p := heap.Pop(m).(Junction)

		ja, hasA := circuits[p.a]

		jb, hasB := circuits[p.b]

		// exists a circuit between these connections already
		if hasA && hasB && ja == jb {
			continue
		}

		// merge the two circuits
		if hasA && hasB && ja != jb {
			for _, j := range jb.boxes {
				circuits[j] = ja
			}
			ja.boxes = append(ja.boxes, jb.boxes...)
			continue
		}

		if !hasA && !hasB {
			l := &Circuit{[]Box{p.a, p.b}}
			circuits[p.a] = l
			circuits[p.b] = l
		} else if !hasA {
			jb.boxes = append(jb.boxes, p.a)
			circuits[p.a] = jb
			circuits[p.b] = jb
		} else {
			ja.boxes = append(ja.boxes, p.b)
			circuits[p.a] = ja
			circuits[p.b] = ja
		}
	}

	x := make(map[*Circuit]bool)

	intHeap := adventheap.New[int](func(a, b int) bool {
		return a-b > 0
	})

	for _, circuit := range circuits {
		if _, has := x[circuit]; has {
			continue
		}
		heap.Push(intHeap, len(circuit.boxes))
		x[circuit] = true
	}

	total := 1

	for i := 0; i < 3; i++ {
		total *= heap.Pop(intHeap).(int)
	}

	fmt.Println(total)
}
