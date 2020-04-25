package modem

import (
	"testing"

	"github.com/richsoap/ldpc/util"
)

func TestModem(t *testing.T) {
	inputData := []int{0, 1, 1, 1, 0, 0}
	neededData := []float64{1, -1, -1, -1, 1, 1}
	mod := Modem{}
	output := mod.Modulate(inputData)
	if !util.FloatSliceCompare(neededData, output) {
		t.Error("Modelate error")
	}
	demodeResult := mod.Demodulate(output)
	if !util.IntSliceCompare(inputData, demodeResult) {
		t.Error("Demodelate error")
	}
}
