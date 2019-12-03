package main

import "fmt"

// Puzzle input : https://adventofcode.com/2019/day/1/input
func main() {
	total := 0
	for {
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil || i < 0 {
			fmt.Println("Please enter a valid integer")
			if i == -100 {
				break
			}
			continue
		}
		fmt.Printf("mass = %d\n", i)
		total = total + fuelReq(i)
		fmt.Printf("total=%d\n", total)
	}
}

func fuelReq(input int) int {
	total := 0
	for {
		div := input / 3
		if div < 3 {
			return total
		}
		fInput := div - 2
		total = total + fInput
		input = fInput
	}
	return total
}
