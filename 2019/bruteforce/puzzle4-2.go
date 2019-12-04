package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	for i := 158126; i < 624574; i++ {
		if isValid(i) {
			count++
		}
	}
	fmt.Printf("Final result %d", count)
}

func isValid(num int) bool {
	a := intToArray(num)
	for i := 1; i <= 5; i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	repeat := map[int]int{}
	for _, r := range a {
		repeat[r]++
	}
	for _, r := range repeat {
		if r == 2 {
			return true
		}
	}
	return false
}

func intToArray(num int) [6]int {
	a := [6]int{}
	for i := 5; i >= 0; i-- {
		a[5-i] = num / int(math.Pow10(i))
		num = num % int(math.Pow10(i))
	}
	return a
}
