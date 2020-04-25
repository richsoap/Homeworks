package coder

import (
	"github.com/richsoap/ldpc/matrix"
)

type Encoder struct {
	HMatrix   matrix.BinaryMatrix
	HsTMatrix matrix.BinaryMatrix
	W, H, Z   int
}

func MakeEncoder(data [][]int, h, w, z int) Encoder {
	result := Encoder{}
	result.HMatrix = matrix.MakeBinaryMatrixWithData(data)
	result.HsTMatrix = matrix.MakeBinaryMatrix(w-h, h)
	result.W = w
	result.H = h
	result.Z = z
	HsData := make([][]int, len(data), len(data))
	for i := range data {
		HsData[i] = data[i][w-h:]
	}
	result.HsTMatrix = matrix.MakeBinaryMatrixWithData(HsData)
	result.HsTMatrix = result.HsTMatrix.Trans()
	return result
}

func (e *Encoder) indexTrans(i int) int {
	return i*e.Z%e.H + i/(e.H/e.Z)
}

func (e *Encoder) Encode(data []int) ([]int, error) {
	p := make([]int, e.H, e.H)
	ts := [][]int{data}
	s := matrix.MakeBinaryMatrixWithData(ts)
	tw, err := s.Mul(&e.HsTMatrix)
	if err != nil {
		return nil, err
	}
	w, err := tw.GetLine(0)
	if err != nil {
		return nil, err
	}
	prevp := 0
	for i := 0; i < e.H; i++ {
		curr := e.indexTrans(i)
		p[curr] = prevp ^ w[curr]
		prevp = p[curr]
	}
	return append(p, data...), nil
}
