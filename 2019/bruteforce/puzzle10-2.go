package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	_, bestCoordinate := findBest()
	fmt.Printf("Final Answer - Puzzle2 : %v", vaporizeAsteroidsUntil(bestCoordinate, 200))
}

func findBest() (int, coordinate) {
	best := 0
	var bestCoordinate coordinate
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
			bestCoordinate = curr
		}
	}
	return best, bestCoordinate
}

func vaporizeAsteroidsUntil(from coordinate, until int) coordinate {
	angles := []float64{}
	vaporizeMap := make(map[float64]map[int]coordinate)
	for _, aLoc := range astroidMap {
		if from == aLoc {
			continue
		}
		angle := from.angle(aLoc)
		d := from.distance(aLoc)
		if _, ok := vaporizeMap[angle]; !ok {
			angles = append(angles, angle)
			m := make(map[int]coordinate)
			vaporizeMap[angle] = m
		}
		m := vaporizeMap[angle]
		m[d] = aLoc
	}
	sort.Float64s(angles)
	vaporized := 0
	for {
		for _, angle := range angles {
			m := vaporizeMap[angle]
			c := 5000
			for d, _ := range m {
				if d < c {
					c = d
				}
			}
			vaporized++
			if vaporized == until {
				return m[c]
			}
			delete(m, c)
		}
		newAngles := []float64{}
		for _, angle := range angles {
			m := vaporizeMap[angle]
			if len(m) > 0 {
				newAngles = append(newAngles, angle)
			}
		}
		angles = newAngles
	}
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

func (c coordinate) angle(d coordinate) float64 {
	a := math.Atan2(float64(d.x-c.x), float64(c.y-d.y)) * 180 / math.Pi
	if a < 0 {
		a = 360 + a
	}
	return a
}
