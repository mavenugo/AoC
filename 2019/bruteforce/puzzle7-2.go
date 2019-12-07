package main

import (
	"fmt"
	"strconv"
)

var pInput [523]int

func main() {
	// Input from https://adventofcode.com/2019/day/7/input
	pInput = [523]int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 30, 55, 80, 101, 118, 199, 280, 361, 442, 99999, 3, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 1002, 9, 4, 9, 101, 4, 9, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 5, 9, 9, 1002, 9, 2, 9, 101, 3, 9, 9, 102, 4, 9, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 101, 5, 9, 9, 102, 3, 9, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}
	combinations := []string{}
	i := 0
	Perm([]rune("56789"), func(a []rune) {
		combinations = append(combinations, string(a))
		i++
	})
	result := 0
	for i := 0; i < len(combinations); i++ {
		r := thrust(combinations[i])
		if r > result {
			result = r
		}
	}
	fmt.Printf("Final Answer : %d\n", result)
}

func thrust(combination string) int {
	input := 0
	var chans [5]chan int
	for i := range chans {
		chans[i] = make(chan int)
	}
	fin := make(chan int)
	for i := 0; i < len(combination); i++ {
		r := []rune(combination)[i]
		phase, _ := strconv.Atoi(string([]rune{r}))
		var rcv chan int
		if i == 0 {
			rcv = chans[len(combination)-1]
		} else {
			rcv = chans[i-1]
		}
		last := false
		if i == len(combination)-1 {
			last = true
		}
		go amp(phase, i, last, rcv, chans[i], fin)
	}
	input = <-fin
	return input
}

func amp(phase int, index int, lastAmp bool, rcv chan int, send chan int, fin chan int) int {
	arr := pInput
	args := []int{phase}
	if index == 0 {
		args = append(args, index)
	}
	result := IntCode(args, lastAmp, rcv, send, fin, arr[:])
	if !isClosed(send) {
		close(send)
	}
	if !isClosed(rcv) {
		close(rcv)
	}
	return result
}

func IntCode(args []int, lastAmp bool, rcv chan int, send chan int, fin chan int, input []int) int {
	idx := 0
	iArr := []int{}
	var answer int
	arg := 0
done:
	for {
		var num int
		var next int
		inst := input[idx] % 100
		iArr = append(iArr, inst)
		pm1 := (input[idx] / 100) % 10
		pm2 := (input[idx] / 1000) % 10
		pm3 := (input[idx] / 10000) % 10
		switch inst {
		case 1:
			p1 := input[idx+1]
			if pm1 == 1 {
				p1 = idx + 1
			}
			p2 := input[idx+2]
			if pm2 == 1 {
				p2 = idx + 2
			}
			num = input[p1] + input[p2]
			p3 := input[idx+3]
			if pm3 == 1 {
				p3 = idx + 3
			}
			input[p3] = num
			next = 4
		case 2:
			p1 := input[idx+1]
			if pm1 == 1 {
				p1 = idx + 1
			}
			p2 := input[idx+2]
			if pm2 == 1 {
				p2 = idx + 2
			}
			num = input[p1] * input[p2]
			p3 := input[idx+3]
			if pm3 == 1 {
				p3 = idx + 3
			}
			input[p3] = num
			next = 4
		case 3:
			if arg < len(args) {
				input[input[idx+1]] = args[arg]
				arg++
			} else {
				select {
				case r := <-rcv:
					input[input[idx+1]] = r
				case <-fin:
					return answer
				}

			}
			next = 2
		case 4:
			answer = input[input[idx+1]]
			if !isClosed(send) {
				send <- input[input[idx+1]]
			}
			next = 2
		case 5:
			p1 := input[idx+1]
			if pm1 == 1 {
				p1 = idx + 1
			}
			p2 := input[idx+2]
			if pm2 == 1 {
				p2 = idx + 2
			}
			if input[p1] != 0 {
				idx = input[p2]
				next = 0
			} else {
				next = 3
			}
		case 6:
			p1 := input[idx+1]
			if pm1 == 1 {
				p1 = idx + 1
			}
			p2 := input[idx+2]
			if pm2 == 1 {
				p2 = idx + 2
			}
			if input[p1] == 0 {
				idx = input[p2]
				next = 0
			} else {
				next = 3
			}
		case 7:
			p1 := input[idx+1]
			if pm1 == 1 {
				p1 = idx + 1
			}
			p2 := input[idx+2]
			if pm2 == 1 {
				p2 = idx + 2
			}
			p3 := input[idx+3]
			if pm3 == 1 {
				p3 = idx + 3
			}
			if input[p1] < input[p2] {
				input[p3] = 1
			} else {
				input[p3] = 0
			}
			next = 4
		case 8:
			p1 := input[idx+1]
			if pm1 == 1 {
				p1 = idx + 1
			}
			p2 := input[idx+2]
			if pm2 == 1 {
				p2 = idx + 2
			}
			p3 := input[idx+3]
			if pm3 == 1 {
				p3 = idx + 3
			}
			if input[p1] == input[p2] {
				input[p3] = 1
			} else {
				input[p3] = 0
			}
			next = 4
		case 99:
			break done
		default:
			fmt.Printf("INPUT ERROR %v %v %d %d\n", iArr, input, inst, idx)
		}
		idx = idx + next
	}
	if lastAmp {
		fin <- answer
	}
	return answer
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func isClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}
