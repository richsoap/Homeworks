package coder

import (
	"github.com/richsoap/ldpc/matrix"
)

type Decoder struct {
	HMatrix matrix.BinaryMatrix
	W, H    int
}

func MakeDecoder(data [][]int, h, w int) Decoder {
	result := Decoder{}
	result.HMatrix = matrix.MakeBinaryMatrixWithData(data)
	result.W = w
	result.H = h
	return result
}

func (d *Decoder) Check(data []int) []int {
	tData := [][]int{data}
	x := matrix.MakeBinaryMatrixWithData(tData)
	xT := x.Trans()
	checkResult, _ := d.HMatrix.Mul(&xT)
	result := make([]int, 0)
	for i := 0; i < checkResult.H; i++ {
		if checkResult.Row[i].Head.Next != nil {
			result = append(result, i)
		}
	}
	return result
}
