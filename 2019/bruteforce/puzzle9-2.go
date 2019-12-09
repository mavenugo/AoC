package main

import (
	"fmt"
)

var pInput []int64

func main() {
	// Input from https://adventofcode.com/2019/day/9/input
	pInput = []int64{1102, 34463338, 34463338, 63, 1007, 63, 34463338, 63, 1005, 63, 53, 1101, 3, 0, 1000, 109, 988, 209, 12, 9, 1000, 209, 6, 209, 3, 203, 0, 1008, 1000, 1, 63, 1005, 63, 65, 1008, 1000, 2, 63, 1005, 63, 904, 1008, 1000, 0, 63, 1005, 63, 58, 4, 25, 104, 0, 99, 4, 0, 104, 0, 99, 4, 17, 104, 0, 99, 0, 0, 1102, 1, 37, 1000, 1101, 856, 0, 1029, 1101, 286, 0, 1025, 1101, 39, 0, 1004, 1101, 861, 0, 1028, 1101, 845, 0, 1026, 1102, 28, 1, 1002, 1102, 1, 0, 1020, 1101, 0, 892, 1023, 1101, 0, 291, 1024, 1101, 35, 0, 1018, 1101, 0, 27, 1006, 1102, 1, 26, 1011, 1101, 33, 0, 1019, 1102, 31, 1, 1014, 1102, 1, 36, 1010, 1102, 23, 1, 1007, 1101, 0, 32, 1016, 1101, 29, 0, 1008, 1101, 20, 0, 1001, 1102, 1, 25, 1015, 1101, 38, 0, 1017, 1101, 0, 24, 1012, 1102, 1, 22, 1005, 1101, 1, 0, 1021, 1101, 0, 21, 1003, 1102, 1, 838, 1027, 1102, 1, 30, 1013, 1101, 895, 0, 1022, 1101, 0, 34, 1009, 109, 7, 1208, 0, 22, 63, 1005, 63, 201, 1001, 64, 1, 64, 1105, 1, 203, 4, 187, 1002, 64, 2, 64, 109, -6, 2102, 1, 5, 63, 1008, 63, 24, 63, 1005, 63, 223, 1105, 1, 229, 4, 209, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 17, 21102, 40, 1, -6, 1008, 1012, 40, 63, 1005, 63, 255, 4, 235, 1001, 64, 1, 64, 1106, 0, 255, 1002, 64, 2, 64, 109, -15, 21108, 41, 41, 9, 1005, 1012, 277, 4, 261, 1001, 64, 1, 64, 1106, 0, 277, 1002, 64, 2, 64, 109, 11, 2105, 1, 10, 4, 283, 1105, 1, 295, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -9, 21101, 42, 0, 8, 1008, 1013, 44, 63, 1005, 63, 315, 1105, 1, 321, 4, 301, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 13, 1206, 3, 337, 1001, 64, 1, 64, 1106, 0, 339, 4, 327, 1002, 64, 2, 64, 109, -10, 1208, 0, 29, 63, 1005, 63, 361, 4, 345, 1001, 64, 1, 64, 1106, 0, 361, 1002, 64, 2, 64, 109, 2, 2108, 27, -4, 63, 1005, 63, 383, 4, 367, 1001, 64, 1, 64, 1105, 1, 383, 1002, 64, 2, 64, 109, -4, 1207, 2, 30, 63, 1005, 63, 405, 4, 389, 1001, 64, 1, 64, 1105, 1, 405, 1002, 64, 2, 64, 109, 22, 1205, -8, 417, 1106, 0, 423, 4, 411, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -27, 2108, 19, 0, 63, 1005, 63, 443, 1001, 64, 1, 64, 1106, 0, 445, 4, 429, 1002, 64, 2, 64, 109, 13, 21108, 43, 45, -1, 1005, 1013, 461, 1106, 0, 467, 4, 451, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 1, 21107, 44, 45, 4, 1005, 1019, 485, 4, 473, 1105, 1, 489, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -8, 2102, 1, -7, 63, 1008, 63, 37, 63, 1005, 63, 515, 4, 495, 1001, 64, 1, 64, 1106, 0, 515, 1002, 64, 2, 64, 109, 1, 2107, 38, -4, 63, 1005, 63, 533, 4, 521, 1105, 1, 537, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 4, 21107, 45, 44, 1, 1005, 1013, 553, 1106, 0, 559, 4, 543, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -7, 2107, 21, -4, 63, 1005, 63, 575, 1106, 0, 581, 4, 565, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 9, 1205, 7, 599, 4, 587, 1001, 64, 1, 64, 1105, 1, 599, 1002, 64, 2, 64, 109, -11, 2101, 0, -3, 63, 1008, 63, 40, 63, 1005, 63, 619, 1105, 1, 625, 4, 605, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 1, 2101, 0, -2, 63, 1008, 63, 28, 63, 1005, 63, 651, 4, 631, 1001, 64, 1, 64, 1106, 0, 651, 1002, 64, 2, 64, 109, 1, 21102, 46, 1, 7, 1008, 1012, 44, 63, 1005, 63, 671, 1106, 0, 677, 4, 657, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 4, 1201, -7, 0, 63, 1008, 63, 28, 63, 1005, 63, 699, 4, 683, 1105, 1, 703, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -6, 1207, -3, 36, 63, 1005, 63, 719, 1105, 1, 725, 4, 709, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -4, 1201, 6, 0, 63, 1008, 63, 23, 63, 1005, 63, 745, 1106, 0, 751, 4, 731, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 8, 1202, -6, 1, 63, 1008, 63, 20, 63, 1005, 63, 777, 4, 757, 1001, 64, 1, 64, 1105, 1, 777, 1002, 64, 2, 64, 109, 5, 1202, -5, 1, 63, 1008, 63, 25, 63, 1005, 63, 801, 1001, 64, 1, 64, 1105, 1, 803, 4, 783, 1002, 64, 2, 64, 109, 8, 21101, 47, 0, -6, 1008, 1014, 47, 63, 1005, 63, 829, 4, 809, 1001, 64, 1, 64, 1106, 0, 829, 1002, 64, 2, 64, 109, 1, 2106, 0, 6, 1001, 64, 1, 64, 1106, 0, 847, 4, 835, 1002, 64, 2, 64, 109, 11, 2106, 0, -4, 4, 853, 1105, 1, 865, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -15, 1206, 3, 883, 4, 871, 1001, 64, 1, 64, 1106, 0, 883, 1002, 64, 2, 64, 109, 14, 2105, 1, -8, 1105, 1, 901, 4, 889, 1001, 64, 1, 64, 4, 64, 99, 21102, 1, 27, 1, 21102, 1, 915, 0, 1106, 0, 922, 21201, 1, 57564, 1, 204, 1, 99, 109, 3, 1207, -2, 3, 63, 1005, 63, 964, 21201, -2, -1, 1, 21102, 1, 942, 0, 1105, 1, 922, 22101, 0, 1, -1, 21201, -2, -3, 1, 21101, 957, 0, 0, 1105, 1, 922, 22201, 1, -1, -2, 1106, 0, 968, 21202, -2, 1, -2, 109, -3, 2106, 0, 0}
	//pInput = []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	//pInput = []int64{104, 1125899906842624, 99}
	//pInput = []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	fmt.Printf("Final Answer : %v\n", boost())
}

func boost() int64 {
	args := []int{2}
	return IntCode(args, true, nil, nil, nil, pInput)
}

func IntCode(args []int, lastAmp bool, rcv chan int64, send chan int64, fin chan int64, input []int64) int64 {
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
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			p2 := getInput(idx + 2)
			if pm2 == 1 {
				p2 = idx + 2
			} else if pm2 == 2 {
				p2 = (relativeBase + getInput(idx+2))
			}
			num = getInput(p1) + getInput(p2)
			p3 := getInput(idx + 3)
			if pm3 == 1 {
				p3 = idx + 3
			} else if pm3 == 2 {
				p3 = (relativeBase + getInput(idx+3))
			}
			setInput(p3, num)
			next = 4
		case 2:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			p2 := getInput(idx + 2)
			if pm2 == 1 {
				p2 = idx + 2
			} else if pm2 == 2 {
				p2 = (relativeBase + getInput(idx+2))
			}
			num = getInput(p1) * getInput(p2)
			p3 := getInput(idx + 3)
			if pm3 == 1 {
				p3 = idx + 3
			} else if pm3 == 2 {
				p3 = (relativeBase + getInput(idx+3))
			}
			setInput(p3, num)
			next = 4
		case 3:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			if arg < len(args) {
				setInput(p1, int64(args[arg]))
				arg++
			} else {
				select {
				case r := <-rcv:
					setInput(p1, r)
				case <-fin:
					return answer
				}

			}
			next = 2
		case 4:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			answer = getInput(p1)
			fmt.Printf("Answer : %d\n", answer)
			if !isClosed(send) {
				send <- getInput(p1)
			}
			next = 2
		case 5:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			p2 := getInput(idx + 2)
			if pm2 == 1 {
				p2 = idx + 2
			} else if pm2 == 2 {
				p2 = (relativeBase + getInput(idx+2))
			}
			if getInput(p1) != 0 {
				idx = getInput(p2)
				next = 0
			} else {
				next = 3
			}
		case 6:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			p2 := getInput(idx + 2)
			if pm2 == 1 {
				p2 = idx + 2
			} else if pm2 == 2 {
				p2 = (relativeBase + getInput(idx+2))
			}
			if getInput(p1) == 0 {
				idx = getInput(p2)
				next = 0
			} else {
				next = 3
			}
		case 7:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			p2 := getInput(idx + 2)
			if pm2 == 1 {
				p2 = idx + 2
			} else if pm2 == 2 {
				p2 = (relativeBase + getInput(idx+2))
			}
			p3 := getInput(idx + 3)
			if pm3 == 1 {
				p3 = idx + 3
			} else if pm3 == 2 {
				p3 = (relativeBase + getInput(idx+3))
			}
			if getInput(p1) < getInput(p2) {
				setInput(p3, 1)
			} else {
				setInput(p3, 0)
			}
			next = 4
		case 8:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
			p2 := getInput(idx + 2)
			if pm2 == 1 {
				p2 = idx + 2
			} else if pm2 == 2 {
				p2 = (relativeBase + getInput(idx+2))
			}
			p3 := getInput(idx + 3)
			if pm3 == 1 {
				p3 = idx + 3
			} else if pm3 == 2 {
				p3 = (relativeBase + getInput(idx+3))
			}
			if getInput(p1) == getInput(p2) {
				setInput(p3, 1)
			} else {
				setInput(p3, 0)
			}
			next = 4
		case 9:
			p1 := getInput(idx + 1)
			if pm1 == 1 {
				p1 = idx + 1
			} else if pm1 == 2 {
				p1 = (relativeBase + getInput(idx+1))
			}
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
