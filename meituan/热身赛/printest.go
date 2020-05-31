package main

import (
	"fmt"
	"math"
)

func main() {
	//inputs := []float64{6.347, 0.099, 5.673}
	inputs := []float64{1, 1, -1}
	product := 1.0
	for i := range inputs {
		product *= math.Tanh(inputs[i] / 2)
	}
	fmt.Printf("%v", 2*math.Atanh(product))
}
