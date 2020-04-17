package coder

import (
	"fmt"
	"testing"

	"github.com/richsoap/ldpc/source"
	"github.com/richsoap/ldpc/util"
)

func TestCode(t *testing.T) {
	HMatrix, z, err := util.LoadHMatrix("../doc/H_block.txt")
	if err != nil {
		t.Error(err)
		return
	}
	sigLen := len(HMatrix)
	encoder := MakeEncoder(HMatrix, len(HMatrix), len(HMatrix[0]), z)
	decoder := MakeDecoder(HMatrix, len(HMatrix), len(HMatrix[0]))
	sigSrc := source.Source{}
	inputData := sigSrc.Create(sigLen)
	codeResult, err := encoder.Encode(inputData)
	if err != nil {
		t.Error("Encode Error Occurs")
		t.Error(err)
		return
	}
	checkResult := decoder.Check(codeResult)
	if len(checkResult) != 0 {
		t.Error(fmt.Sprintf("Encode Check Error, error bits %v: %v", len(checkResult), checkResult))
		if checkResult[0] == 0 {
			t.Error(fmt.Sprintf("p[0]: %v", codeResult[0]))
			checkBit := 0
			for i := range inputData {
				if res, _ := decoder.HMatrix.Get(0, i+sigLen); res == 1 && inputData[i] == 1 {
					checkBit++
				}
			}
			t.Error(fmt.Sprintf("need: %v", checkBit%2))
		}
	}
}
