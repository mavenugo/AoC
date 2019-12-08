package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("puzzle6.input")
	scanner := bufio.NewScanner(f)
	ro := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		m := strings.Split(line, ")")
		ro[m[1]] = m[0]
	}

	fmt.Printf("Final Answer : %d\n", calculateTransfers(ro, "YOU", "SAN"))
}

func calculateTransfers(ro map[string]string, from string, to string) int {
	fOrbits := ll(ro, from)
	tOrbits := ll(ro, to)
	fNum := 0
	tNum := 0
outer:
	for _, f := range fOrbits {
		for _, t := range tOrbits {
			if f == t {
				break outer
			}
			tNum++
		}
		fNum++
		tNum = 0
	}
	return fNum + tNum
}

func ll(ro map[string]string, key string) []string {
	v, ok := ro[key]
	if !ok {
		return []string{}
	}
	orbits := []string{v}
	o := ll(ro, v)
	orbits = append(orbits, o...)
	return orbits
}
