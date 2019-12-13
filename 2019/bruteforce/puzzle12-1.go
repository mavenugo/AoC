package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("puzzle12.input")
	scanner := bufio.NewScanner(f)
	var moons []*moon

	for scanner.Scan() {
		line := scanner.Text()
		coordinates := strings.Split(line, ",")
		X := strings.Split(string([]rune(coordinates[0])[1:]), "=")
		Y := strings.Split(coordinates[1], "=")
		Z := strings.Split(string([]rune(coordinates[2])[:len(coordinates[2])-1]), "=")
		x, _ := strconv.Atoi(X[1])
		y, _ := strconv.Atoi(Y[1])
		z, _ := strconv.Atoi(Z[1])
		c := coordinate{x, y, z}
		moon := moon{c: c}
		moons = append(moons, &moon)
	}
	fmt.Printf("Final Answer : %v\n", totalEnergy(moons, 1000))
}

func totalEnergy(moons []*moon, steps int) int {
	for s := 0; s < steps; s++ {
		//calculate gravity
		for i, target := range moons {
			for j, moon := range moons {
				if i == j {
					continue
				}
				if target.c.x > moon.c.x {
					target.v.x--
				} else if target.c.x < moon.c.x {
					target.v.x++
				}
				if target.c.y > moon.c.y {
					target.v.y--
				} else if target.c.y < moon.c.y {
					target.v.y++
				}
				if target.c.z > moon.c.z {
					target.v.z--
				} else if target.c.z < moon.c.z {
					target.v.z++
				}
			}
		}
		//calculate new position
		for _, target := range moons {
			target.c.x = target.c.x + target.v.x
			target.c.y = target.c.y + target.v.y
			target.c.z = target.c.z + target.v.z
		}
	}
	var total int
	for _, target := range moons {
		pot := abs(target.c.x) + abs(target.c.y) + abs(target.c.z)
		kin := abs(target.v.x) + abs(target.v.y) + abs(target.v.z)
		total += pot * kin
	}
	return total
}

type moon struct {
	c coordinate
	v coordinate
}

type coordinate struct {
	x, y, z int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
