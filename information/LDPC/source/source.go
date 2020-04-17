package source

import (
	"math/rand"
	"time"
)

type Source struct{}

func (s *Source) Create(l int) []int {
	result := make([]int, l, l)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = rand.Intn(2)
	}
	return result
}
