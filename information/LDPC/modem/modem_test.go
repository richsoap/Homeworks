package modem

import (
	"testing"

	"github.com/richsoap/ldpc/util"
)

func TestModem(t *testing.T) {
	inputData := []int{0, 1, 1, 1, 0, 0}
	neededData := []float64{1, -1, -1, -1, 1, 1}
	output := Modulate(inputData)
	if !util.FloatSliceCompare(neededData, output) {
		t.Error("Modelate error")
	}
	demodeResult := Demodulate(output)
	if !util.IntSliceCompare(inputData, demodeResult) {
		t.Error("Demodelate error")
	}
}
