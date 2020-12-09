package main

import(
	"fmt"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
		return false
	}
	record := make([][]bool, len(s1) + 1, len(s1) + 1)
	for index, _ := range record {
		record[index] = make([]bool, len(s2) + 1, len(s2) + 1)
	}
	record[0][0] = true
	for i := 1;i <= len(s1);i ++ {
		record[i][0] = record[i - 1][0] && s1[i - 1] == s3[i - 1]
	}
	for i := 1;i <= len(s2);i ++ {
		record[0][i] = record[0][i - 1] && s2[i - 1] == s3[i - 1]
	}
	for i := 1;i <= len(s1);i ++ {
		for j := 1;j <= len(s2);j ++ {
			record[i][j] = (record[i - 1][j] && s1[i - 1] == s3[i + j - 1]) || (record[i][j - 1] && s2[j - 1] == s3[i + j - 1])
		}
	}
	return record[len(s1)][len(s2)]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))
}