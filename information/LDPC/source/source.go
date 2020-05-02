package source

import (
	"math/rand"
	"time"
)

func Create(l int) []int {
	result := make([]int, l, l)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = rand.Intn(2)
	}
	return result
}
