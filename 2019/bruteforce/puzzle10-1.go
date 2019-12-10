package main

import (
	"bufio"
	"fmt"
	"os"
)

var astroidMap [][]rune

var EMPTY = rune(46)
var LOCATED = rune(35)
var totalAstroids int
var totalEmpty int

func main() {
	f, _ := os.Open("puzzle10.input")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		mapAstroid(line)
	}
	for _, aArray := range astroidMap {
		for _, a := range aArray {
			if a == LOCATED {
				totalAstroids++
			} else {
				totalEmpty++
			}
		}
	}
	fmt.Printf("Final Answer : %d", findBestNumbers())
}

func findBestNumbers() int {
	best := 0
	for x, aArray := range astroidMap {
		for y, a := range aArray {
			if a == LOCATED {
				n := numAstroidsInLOS(x, y)
				if n > best {
					best = n
				}
			}
		}
	}
	return best
}

func numAstroidsInLOS(x int, y int) int {
	numBlocked := 0
	blocked := false
	//EAST
	for e := x + 1; e < len(astroidMap); e++ {
		if astroidMap[e][y] == LOCATED {
			if !blocked {
				blocked = true
			} else {
				numBlocked++
			}
		}
	}
	blocked = false
	//WEST
	for w := x - 1; w >= 0; w-- {
		if astroidMap[w][y] == LOCATED {
			if !blocked {
				blocked = true
			} else {
				numBlocked++
			}
		}
	}
	blocked = false
	//NORTH
	for n := y - 1; n >= 0; n-- {
		if astroidMap[x][n] == LOCATED {
			if !blocked {
				blocked = true
			} else {
				numBlocked++
			}
		}
	}
	blocked = false
	//SOUTH
	for s := y + 1; s < len(astroidMap[x]); s++ {
		if astroidMap[x][s] == LOCATED {
			if !blocked {
				blocked = true
			} else {
				numBlocked++
			}
		}
	}

	var tx int
	//NE
	//Immediate East row wont be blocking ... hence starting with x+2
	for e := x + 2; e < len(astroidMap); e++ {
		//Immediate North row wont be blocking ... hence starting with y-2
		for n := y - 2; n >= 0; n-- {
			if astroidMap[e][n] == LOCATED {
				if (e-x)%(y-n) == 0 {
					tx = e
					for s := n + 1; s >= y-1; s++ {
						tx = tx - (e-x)/(y-n)
						if astroidMap[tx][s] == LOCATED {
							numBlocked++
							break
						}

					}

				}
			}
		}
	}

	//NW
	//Immediate West row wont be blocking ... hence starting with x-2
	for w := x - 2; w >= 0; w-- {
		//Immediate North row wont be blocking ... hence starting with y-2
		for n := y - 2; n >= 0; n-- {
			if astroidMap[w][n] == LOCATED {
				if (y-n)%(x-w) == 0 {
					tx = w
					for s := n + 1; s >= y-1; s++ {
						tx = tx + (y-n)/(x-w)
						if astroidMap[tx][s] == LOCATED {
							numBlocked++
							break
						}

					}

				}
			}
		}
	}

	//SE
	//Immediate East row wont be blocking ... hence starting with x+2
	for e := x + 2; e < len(astroidMap); e++ {
		//Immediate South row wont be blocking ... hence starting with y+2
		for s := y + 2; s < len(astroidMap[e]); s++ {
			if astroidMap[e][s] == LOCATED {
				if e%s == 0 {
					tx = e
					for n := s - 1; n < y-1; n-- {
						tx = tx - e/s
						if astroidMap[tx][n] == LOCATED {
							numBlocked++
							break
						}

					}

				}
			}
		}
	}
	//SW
	//Immediate West row wont be blocking ... hence starting with x+2
	for w := x - 2; w >= 0; w-- {
		//Immediate South row wont be blocking ... hence starting with y+2
		for s := y + 2; s < len(astroidMap[w]); s++ {
			if astroidMap[w][s] == LOCATED {
				if (s-y)%(x-w) == 0 {
					tx = w
					for n := s - 1; n < y-1; n-- {
						tx = tx + (s-y)/(x-w)
						if astroidMap[tx][n] == LOCATED {
							numBlocked++
							break
						}

					}

				}
			}
		}
	}
	fmt.Printf("(%d - %d = %d) ", totalAstroids, numBlocked, totalAstroids-numBlocked)
	return totalAstroids - numBlocked
}

func mapAstroid(coordinates string) {
	astroidMap = append(astroidMap, []rune(coordinates))
}
