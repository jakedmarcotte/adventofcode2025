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
	position := 50
	password_on_zero := 0
	password_passing_zero := 0
	for scanner.Scan() {

		/**
		For each turn:
		- full rotation: distance / 100
		- pass over 0:
		- final position of 0
		*/
		turn := scanner.Text()
		direction := turn[0]
		distance := 0
		full_rotations := 0
		fmt.Sscanf(turn[1:], "%d", &distance)
		og_distance := distance
		og_position := position
		operator := '+'

		if distance > 99 {
			full_rotations += distance / 100
			distance = distance % 100
		}

		if direction == 'R' {
			position += distance
			if position > 99 {
				position = position % 100
				full_rotations++
			}
		} else if direction == 'L' {
			position -= distance
			operator = '-'
			if position < 0 {
				position = 100 + position
				if og_position != 0 {
					full_rotations++
				}
			} else if position == 0 {
				full_rotations++
			}
		}

		if og_position != 0 && position == 0 {
			password_on_zero++
		}

		password_passing_zero += full_rotations
		fmt.Printf("P: %d %c %d = %d, Password: %d Password on 0: %d\n", og_position, operator, og_distance, position, password_passing_zero, password_on_zero)

	}
}
