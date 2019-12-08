package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("puzzle8.input")
	scanner := bufio.NewScanner(f)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	imgSize := 25 * 6
	layers := SplitSubN(line, imgSize)
	zeroes := imgSize
	var layer string
	fmt.Printf("line_length: %d layers# : %d\n", len(line), len(layers))
	for _, l := range layers {
		p := make(map[string]int)
		for _, r := range []rune(l) {
			n, _ := p[string(r)]
			p[string(r)] = n + 1
		}
		fmt.Printf("%v\n", p)
		if zeroes > p["0"] {
			zeroes = p["0"]
			layer = l
		}
	}

	p := make(map[string]int)
	for _, r := range []rune(layer) {
		n, _ := p[string(r)]
		p[string(r)] = n + 1
	}
	ones := p["1"]
	twos := p["2"]
	fmt.Printf("Final Answer : %d", ones*twos)
}

func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
