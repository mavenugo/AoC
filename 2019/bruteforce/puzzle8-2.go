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
	var final [25 * 6]string
	for i, _ := range final {
		final[i] = "2"
	}
	for _, l := range layers {
		for i, r := range []rune(l) {
			if final[i] != "2" {
				continue
			}
			if string(r) == "1" {
				final[i] = "#"
			} else if string(r) == "0" {
				final[i] = " "
			}
		}
	}

	for i, _ := range final {
		if i%25 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%s", final[i])
	}
	fmt.Println("")
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
