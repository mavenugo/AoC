package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("puzzle6-1.input")
	scanner := bufio.NewScanner(f)
	ro := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		m := strings.Split(line, ")")
		ro[m[1]] = m[0]
	}

	count := 0
	for k, _ := range ro {
		count += calculateOrbits(ro, k)
	}
	fmt.Printf("Final Answer : %d", count)
}

func calculateOrbits(ro map[string]string, key string) int {
	v, ok := ro[key]
	if !ok {
		return 0
	}
	i := calculateOrbits(ro, v)
	i++
	return i
}
