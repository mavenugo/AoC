package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var astroidMap []coordinate

var EMPTY = 46   // .
var LOCATED = 35 // #
var totalAstroids int
var totalEmpty int

func main() {
	f, _ := os.Open("puzzle10.input")
	scanner := bufio.NewScanner(f)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, hit := range line {
			if hit == rune(LOCATED) {
				astroidMap = append(astroidMap, coordinate{x, y})
			}
		}
		y++
	}

	fmt.Printf("Final Answer : %d", findBestNumbers())
}

func findBestNumbers() int {
	best := 0
	for _, curr := range astroidMap {
		angles := make(map[float64]coordinate)
		c := 0
		for _, running := range astroidMap {
			if curr != running {
				angle := curr.angle(running)
				if _, ok := angles[angle]; !ok {
					angles[angle] = running
					c++
				}
			}
		}
		if c > best {
			best = c
		}
	}
	return best
}

type coordinate struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c coordinate) distance(t coordinate) int {
	return abs(c.x-t.x) + abs(c.y-t.y)
}

func (c coordinate) angle(t coordinate) float64 {
	a := math.Atan2(float64(t.x-c.x), float64(c.y-t.y)) * 180 / math.Pi
	if a < 0 {
		a = 360 + a
	}
	return a
}
