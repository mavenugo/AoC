package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	moons := initialMoons()
	fmt.Printf("Final Answer : %v\n", previousStateMatch(moons))
}

func initialMoons() []*moon {
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
	return moons
}

func previousStateMatch(moons []*moon) int {
	is := initialMoons()
	axis := make(map[string]int)
	var xdone, ydone, zdone bool
	for {
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

		//match on individual axis
		xmatch := true
		ymatch := true
		zmatch := true
		for i, m := range moons {
			if !xdone && (m.c.x != is[i].c.x || m.v.x != is[i].v.x) {
				xmatch = false
			}
			if !ydone && (m.c.y != is[i].c.y || m.v.y != is[i].v.y) {
				ymatch = false
			}
			if !zdone && (m.c.z != is[i].c.z || m.v.z != is[i].v.z) {
				zmatch = false
			}
		}
		if xmatch {
			xdone = true
		} else {
			axis["x"]++
		}
		if ymatch {
			ydone = true
		} else {
			axis["y"]++
		}
		if zmatch {
			zdone = true
		} else {
			axis["z"]++
		}
		if xmatch && ymatch && zmatch {
			break
		}
	}
	return LCM(axis["x"]+1, axis["y"]+1, axis["z"]+1)
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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
