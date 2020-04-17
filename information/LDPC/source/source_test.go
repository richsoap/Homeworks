package source

import (
	"fmt"
	"math"
	"testing"
)

func checkBalance(t *testing.T, data []int) {
	n := float64(len(data))
	c := 0
	for _, v := range data {
		c += v
	}
	fc := float64(c)
	if math.Abs(0.5-fc/n) > 0.01 {
		t.Error(fmt.Sprintf("Not balance: %v/%v %v", fc, n, fc/n))
	}
}

func TestSource(t *testing.T) {
	source := Source{}
	testResult := source.Create(100000)
	checkBalance(t, testResult)
}
