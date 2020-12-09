package main

import(
	"fmt"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func isValid(s *string) bool {
	if val, err := strconv.Atoi(*s); err != nil || val >= 256 || (len(*s) > 1 && (*s)[0] == '0'){
		return false
	} else {
		return true
	}
}

func getValidAddr(s string, stage int) []string {
	if len(s) == 0 {
		return make([]string, 0, 0)
	}
	if stage == 1 {
		if isValid(&s) {
			return []string{s};
		} else {
			return make([]string, 0, 0)
		}
	}
	result := make([]string, 0)
	workspace := make([]string, 2, 2)
	for i := 1;i < min(4, len(s));i ++ {
		workspace[0] = s[:i]
		if !isValid(&workspace[0]) {
			return result
		}
		tails := getValidAddr(s[i:], stage - 1)
		for _, tail := range tails {
			workspace[1] = tail
			result = append(result, strings.Join(workspace, "."))
		}
	}
	return result
}

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12{
		return make([]string, 0)
	}
	return getValidAddr(s, 4)
}

func main() {
	result := restoreIpAddresses("010010")
	for _, val := range result {
		fmt.Println(val)
	}
}