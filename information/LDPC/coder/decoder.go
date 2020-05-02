package coder

import (
	"sync"

	"github.com/richsoap/ldpc/matrix"
)

type Decoder struct {
	HMatrix    matrix.BinaryMatrix
	W, H       int
	valueNodes []Node
	checkNodes []Node
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

func (d *Decoder) Iterate() {
	var wg sync.WaitGroup
	wg.Add(d.H)
	for i := 0; i < d.H; i++ {
		go d.checkNodes[i].Update(&d.valueNodes, &wg)
	}
	wg.Wait()
	wg.Add(d.W)
	for i := 0; i < d.W; i++ {
		go d.valueNodes[i].Update(&d.checkNodes, &wg)
	}
	wg.Wait()
}

func (d *Decoder) IterateSingle() {
	for i := 0; i < d.H; i++ {
		d.checkNodes[i].Update(&d.valueNodes, nil)
	}
	for i := 0; i < d.W; i++ {
		d.valueNodes[i].Update(&d.checkNodes, nil)
	}
}

func (d *Decoder) LoadData(data []float64, checkFunc int, param ...float64) {
	d.checkNodes = make([]Node, d.H, d.H)
	d.valueNodes = make([]Node, d.W, d.W)

	for index := 0; index < d.H; index++ {
		d.checkNodes[index] = BuildNode(index, 0, BuildRecorder(checkFunc, param...).(Recoder))
		for rol := d.HMatrix.Row[index].Head.Next; rol != nil; rol = rol.Next {
			d.checkNodes[index].AddLink(rol.Val)
		}
	}
	for index := 0; index < d.W; index++ {
		d.valueNodes[index] = BuildNode(index, data[index], BuildRecorder(ValueSum, param...).(Recoder))
		for row := d.HMatrix.Col[index].Head.Next; row != nil; row = row.Next {
			d.valueNodes[index].AddLink(row.Val)
		}
	}
}

func (d *Decoder) GetResult() []float64 {
	result := make([]float64, d.W, d.W)
	for i := 0; i < d.W; i++ {
		result[i] = d.valueNodes[i].GetResult()
	}
	return result
}
