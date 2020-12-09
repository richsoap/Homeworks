package main

import "fmt"

func main() {
	record := [][]int{
		{0, 0, 0, 4, 8, 6, 0, 0, 0},
		{0, 0, 0, 3, 5, 7, 0, 0, 0},
		{0, 0, 0, 2, 1, 9, 0, 0, 0},
		{0, 0, 0, 9, 0, 8, 0, 0, 0},
		{0, 0, 0, 7, 0, 0, 9, 0, 0},
		{0, 0, 0, 6, 0, 0, 0, 8, 0},
		{0, 0, 0, 5, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{5, 0, 4, 8, 0, 0, 0, 0, 0}}
	lines := make([][]int, 6, 6)
	lines[0] = []int{60, 50, 40, 30, 31, 42, 51}
	lines[1] = []int{45, 54, 65, 75}
	lines[2] = []int{87, 86, 76, 66, 56}
	lines[3] = []int{58, 68, 78, 88}
	lines[4] = []int{47, 37}
	lines[5] = []int{28, 27, 26}
	if goDeep(0, 0, &record, &lines) {
		fmt.Printf("%v", record)
	} else {
		fmt.Printf("get nothing")
	}
}

func goDeep(r, c int, record, lines *[][]int) bool {
	if r >= 9 {
		return checkLines(record, lines)
	}
	if (*record)[r][c] != 0 {
		return goDeep(r, c+1, record, lines)
	}
	showed := make(map[int]int)
	for i := 0; i < 9; i++ {
		if i == r {
			continue
		}
		showed[(*record)[i][c]] = 0
	}
	for i := 0; i < 9; i++ {
		if i == c {
			continue
		}
		showed[(*record)[r][i]] = 0
	}
	for i := 1; i <= 9; i++ {
		if _, ok := showed[i]; ok {
			continue
		}
		(*record)[r][c] = i
		if goDeep(r+(c+1)/9, (c+1)%9, record, lines) {
			return true
		}
	}
	(*record)[r][c] = 0
	return false
}

func checkLines(record, lines *[][]int) bool {
	for i := range *lines {
		for j := range (*lines)[i] {
			if j != 0 {
				pos := (*lines)[i][j]
				prev := (*lines)[i][j-1]
				if (*record)[pos/10][pos%10] <= (*record)[prev/10][prev%10] {
					return false
				}
			}
		}
	}
	return true
}
