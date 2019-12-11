package main

import (
	"fmt"
)

const LEFT = 0
const RIGHT = 1
const UP = 2
const DOWN = 3

var pInput []int64
var panels map[coordinate]panel
var licenseArray []coordinate
var currPosition coordinate
var currOutputIdx int
var currDirection = UP

func main() {
	// Input from https://adventofcode.com/2019/day/11/input
	pInput = []int64{3, 8, 1005, 8, 320, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 29, 2, 101, 10, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 54, 2, 3, 16, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 81, 1006, 0, 75, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 105, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 128, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 149, 1, 105, 5, 10, 1, 105, 20, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 179, 1, 101, 1, 10, 2, 109, 8, 10, 1006, 0, 74, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 213, 1006, 0, 60, 2, 1105, 9, 10, 1, 1005, 11, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1002, 8, 1, 245, 1, 6, 20, 10, 1, 1103, 11, 10, 2, 6, 11, 10, 2, 1103, 0, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1002, 8, 1, 284, 2, 1103, 12, 10, 2, 1104, 14, 10, 2, 1004, 12, 10, 2, 1009, 4, 10, 101, 1, 9, 9, 1007, 9, 968, 10, 1005, 10, 15, 99, 109, 642, 104, 0, 104, 1, 21102, 1, 48063419288, 1, 21102, 1, 337, 0, 1105, 1, 441, 21101, 0, 846927340300, 1, 21101, 0, 348, 0, 1105, 1, 441, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21102, 1, 235245104151, 1, 21102, 395, 1, 0, 1105, 1, 441, 21102, 29032123584, 1, 1, 21101, 0, 406, 0, 1105, 1, 441, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 0, 709047878500, 1, 21101, 429, 0, 0, 1106, 0, 441, 21101, 868402070284, 0, 1, 21102, 1, 440, 0, 1105, 1, 441, 99, 109, 2, 22102, 1, -1, 1, 21101, 40, 0, 2, 21101, 0, 472, 3, 21102, 462, 1, 0, 1105, 1, 505, 109, -2, 2106, 0, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 467, 468, 483, 4, 0, 1001, 467, 1, 467, 108, 4, 467, 10, 1006, 10, 499, 1102, 1, 0, 467, 109, -2, 2106, 0, 0, 0, 109, 4, 2101, 0, -1, 504, 1207, -3, 0, 10, 1006, 10, 522, 21101, 0, 0, -3, 22101, 0, -3, 1, 21202, -2, 1, 2, 21101, 1, 0, 3, 21102, 541, 1, 0, 1106, 0, 546, 109, -4, 2106, 0, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 569, 2207, -4, -2, 10, 1006, 10, 569, 21202, -4, 1, -4, 1105, 1, 637, 22102, 1, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 588, 0, 0, 1105, 1, 546, 22102, 1, 1, -4, 21101, 0, 1, -1, 2207, -4, -2, 10, 1006, 10, 607, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 629, 21201, -1, 0, 1, 21102, 629, 1, 0, 106, 0, 504, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}

	panels = make(map[coordinate]panel)
	fmt.Printf("Final Answer : %v\n", paintingRobot())
	normalizedArray := []coordinate{}
	first := true
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	for _, l := range licenseArray {
		if first {
			minX = l.x
			minY = l.y
			first = false
		}
		if l.x < minX {
			minX = l.x
		}
		if l.y < minY {
			minY = l.y
		}
		if l.x > maxX {
			maxX = l.x
		}
		if l.y > maxX {
			maxX = l.y
		}
	}
	for _, l := range licenseArray {
		normalizedArray = append(normalizedArray, coordinate{l.x + minX, l.y + minY})
	}
	var hull [][]int
	hull = make([][]int, maxX+abs(minX)+1)
	for x, _ := range hull {
		hull[x] = make([]int, maxY+abs(minY)+1)
	}

	for c, p := range panels {
		hull[c.x+abs(minX)][c.y+abs(minY)] = p.color
	}

	for x, _ := range hull {
		for y, _ := range hull[x] {
			if hull[x][y] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func paintingRobot() int {
	args := []int{1}
	IntCode(args, true, nil, nil, nil, outputCB, pInput)
	return len(panels)
}

func outputCB(val int) {
	currOutputIdx++
	if currOutputIdx%2 == 1 {
		p := panel{currPosition, val}
		panels[currPosition] = p
		licenseArray = append(licenseArray, currPosition)
	} else {
		//direction
		switch currDirection {
		case UP:
			if val == LEFT {
				currDirection = LEFT
				currPosition = coordinate{currPosition.x - 1, currPosition.y}
			} else {
				currDirection = RIGHT
				currPosition = coordinate{currPosition.x + 1, currPosition.y}
			}
		case DOWN:
			if val == RIGHT {
				currDirection = LEFT
				currPosition = coordinate{currPosition.x - 1, currPosition.y}
			} else {
				currDirection = RIGHT
				currPosition = coordinate{currPosition.x + 1, currPosition.y}
			}
		case LEFT:
			if val == RIGHT {
				currDirection = UP
				currPosition = coordinate{currPosition.x, currPosition.y + 1}
			} else {
				currDirection = DOWN
				currPosition = coordinate{currPosition.x, currPosition.y - 1}
			}
		case RIGHT:
			if val == LEFT {
				currDirection = UP
				currPosition = coordinate{currPosition.x, currPosition.y + 1}
			} else {
				currDirection = DOWN
				currPosition = coordinate{currPosition.x, currPosition.y - 1}
			}
		}
	}
}

func IntCode(args []int, lastAmp bool, rcv chan int64, send chan int64, fin chan int64, cb func(int), input []int64) int64 {
	idx := int64(0)
	iArr := []int{}
	var answer int64
	arg := 0
	relativeBase := int64(0)
done:
	for {
		var num int64
		var next int64
		inst := getInput(idx) % 100
		iArr = append(iArr, int(inst))
		pm1 := (getInput(idx) / 100) % 10
		pm2 := (getInput(idx) / 1000) % 10
		pm3 := (getInput(idx) / 10000) % 10
		switch inst {
		case 1:
			p1 := get(input, idx, pm1, 1, relativeBase)
			p2 := get(input, idx, pm2, 2, relativeBase)
			num = getInput(p1) + getInput(p2)
			p3 := get(input, idx, pm3, 3, relativeBase)
			setInput(p3, num)
			next = 4
		case 2:
			p1 := get(input, idx, pm1, 1, relativeBase)
			p2 := get(input, idx, pm2, 2, relativeBase)
			num = getInput(p1) * getInput(p2)
			p3 := get(input, idx, pm3, 3, relativeBase)
			setInput(p3, num)
			next = 4
		case 3:
			p1 := get(input, idx, pm1, 1, relativeBase)
			if arg < len(args) {
				setInput(p1, int64(args[arg]))
				arg++
			} else {
				if p, ok := panels[currPosition]; ok {
					setInput(p1, int64(p.color))
				} else {
					setInput(p1, 0)
				}
			}
			next = 2
		case 4:
			p1 := get(input, idx, pm1, 1, relativeBase)
			answer = getInput(p1)
			if !isClosed(send) {
				send <- getInput(p1)
			}
			cb(int(answer))
			next = 2
		case 5:
			p1 := get(input, idx, pm1, 1, relativeBase)
			p2 := get(input, idx, pm2, 2, relativeBase)
			if getInput(p1) != 0 {
				idx = getInput(p2)
				next = 0
			} else {
				next = 3
			}
		case 6:
			p1 := get(input, idx, pm1, 1, relativeBase)
			p2 := get(input, idx, pm2, 2, relativeBase)
			if getInput(p1) == 0 {
				idx = getInput(p2)
				next = 0
			} else {
				next = 3
			}
		case 7:
			p1 := get(input, idx, pm1, 1, relativeBase)
			p2 := get(input, idx, pm2, 2, relativeBase)
			p3 := get(input, idx, pm3, 3, relativeBase)
			if getInput(p1) < getInput(p2) {
				setInput(p3, 1)
			} else {
				setInput(p3, 0)
			}
			next = 4
		case 8:
			p1 := get(input, idx, pm1, 1, relativeBase)
			p2 := get(input, idx, pm2, 2, relativeBase)
			p3 := get(input, idx, pm3, 3, relativeBase)
			if getInput(p1) == getInput(p2) {
				setInput(p3, 1)
			} else {
				setInput(p3, 0)
			}
			next = 4
		case 9:
			p1 := get(input, idx, pm1, 1, relativeBase)
			relativeBase = relativeBase + getInput(p1)
			next = 2
		case 99:
			break done
		default:
			fmt.Printf("INPUT ERROR %v %v %d %d\n", iArr, input, inst, idx)
		}
		idx = idx + next
		for i := int64(0); i < idx-int64(len(input)); i++ {
			input = append(input, int64(0))
		}
	}
	if lastAmp && fin != nil {
		fin <- answer
	}
	return answer
}

func getInput(idx int64) int64 {
	extend := (idx - int64(len(pInput))) + 1
	for i := int64(0); i < extend; i++ {
		pInput = append(pInput, int64(0))
	}

	return pInput[idx]

}
func setInput(idx int64, val int64) {
	extend := (idx - int64(len(pInput))) + 1
	for i := int64(0); i < extend; i++ {
		pInput = append(pInput, int64(0))
	}
	pInput[idx] = val
}

func isClosed(ch <-chan int64) bool {
	if ch == nil {
		return true
	}
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func get(input []int64, idx int64, pm int64, ptr int64, relativeBase int64) int64 {
	p := getInput(idx + ptr)
	if pm == 1 {
		p = idx + ptr
	} else if pm == 2 {
		p = (relativeBase + getInput(idx+ptr))
	}
	return p
}

type panel struct {
	c     coordinate
	color int
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
