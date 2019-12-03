package main

import "fmt"

func main() {
	// Input from https://adventofcode.com/2019/day/2/input
	input := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 10, 23, 2, 13, 23, 27, 1, 5, 27, 31, 2, 6, 31, 35, 1, 6, 35, 39, 2, 39, 9, 43, 1, 5, 43, 47, 1, 13, 47, 51, 1, 10, 51, 55, 2, 55, 10, 59, 2, 10, 59, 63, 1, 9, 63, 67, 2, 67, 13, 71, 1, 71, 6, 75, 2, 6, 75, 79, 1, 5, 79, 83, 2, 83, 9, 87, 1, 6, 87, 91, 2, 91, 6, 95, 1, 95, 6, 99, 2, 99, 13, 103, 1, 6, 103, 107, 1, 2, 107, 111, 1, 111, 9, 0, 99, 2, 14, 0, 0}
	// Fix Array to the "1202 program alarm" state
	input[1] = 12
	input[2] = 2
	idx := 0
done:
	for {
		var num int
		switch input[idx] {
		case 1:
			num = input[input[idx+1]] + input[input[idx+2]]
		case 2:
			num = input[input[idx+1]] * input[input[idx+2]]
		case 99:
			break done
		default:
			fmt.Printf("INPUT ERROR %d %d\n", input[idx], idx)
		}
		fmt.Printf("%d->%d->%d\n", idx, input[idx], num)
		input[input[idx+3]] = num
		idx = idx + 4
	}
	fmt.Printf("Final answer : %d\n", input[0])
}
