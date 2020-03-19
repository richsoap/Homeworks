package main

import(
	"fmt"
	"strings"
)

func ceil(a, b int) int {
	if b == 0 {
		return a
	} else if b < 0 {
		return 0
	} else {
		result := a/b
		if a % b != 0 {
			result ++
		}
		return result
	}
}

func leftLine(words []string, width int) string {
	var sb strings.Builder
	used := 0
	for index, word := range words {
		sb.WriteString(word)
		used += (len(word) + 1)
		if index != len(words) - 1 {
			sb.WriteString(" ")
			used ++
		}
	}
	for i := used;i < width;i ++ {
		sb.WriteString(" ")
	}
	return sb.String()
}

func comLine(words []string, width int) string {
	spaceLen := width
	spaceCount := len(words) - 1
	var sb strings.Builder
	for _, word := range words {
		spaceLen -= len(word)
	}
	for _, word := range words {
		sb.WriteString(word)
		spaceNum := ceil(spaceLen, spaceCount)
		for i := 0;i < spaceNum;i ++ {
			sb.WriteString(" ")
		}
		spaceLen -= spaceNum
		spaceCount --
	}
	return sb.String()
}

func fullJustify(words []string, maxWidth int) []string {
	usedCount := 0
	start := 0
	result := make([]string, 0)
	for index, word := range words {
		if usedCount + len(word) > maxWidth {
			result = append(result, comLine(words[start:index], maxWidth))
			start = index
			usedCount = 0
		}
		usedCount += (len(word) + 1)
	}
	if start < len(words) {
		result = append(result, leftLine(words[start:], maxWidth))
	}
	return result
}

func main() {
	words := []string{"What","must","be","acknowledgment","shall","be"}
	maxWidth := 16
	result := fullJustify(words, maxWidth)
	for _, line := range result {
		fmt.Println(line)
	}
}