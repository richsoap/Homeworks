package coder

import (
	"fmt"
	"testing"

	"github.com/richsoap/ldpc/modem"
	"github.com/richsoap/ldpc/source"
	"github.com/richsoap/ldpc/util"
)

func TestEncodeAndDecoder(t *testing.T) {
	HMatrix, z, err := util.LoadHMatrix("../doc/H_block.txt")
	if err != nil {
		t.Error(err)
		return
	}
	sigLen := len(HMatrix)
	encoder := MakeEncoder(HMatrix, len(HMatrix), len(HMatrix[0]), z)
	decoder := MakeDecoder(HMatrix, len(HMatrix), len(HMatrix[0]))
	inputData := source.Create(sigLen)
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
		return
	}
	recorderType := make(map[string]int)
	recorderType["MinSum"] = CheckMinSum
	recorderType["SumProduct"] = CheckSumProduct
	for key, value := range recorderType {
		t.Log(fmt.Sprintf("Type: %v", key))
		// Did Nothing
		BPSKData := modem.Modulate(codeResult)
		decoder.LoadData(BPSKData, CheckMinSum)
		BitResult := modem.Demodulate(decoder.GetResult())
		checkResult = decoder.Check(BitResult)
		if len(checkResult) != 0 {
			t.Error("Did nothing and cannot pass check")
		}
		BitResult[0] = 1 ^ BitResult[0]
		checkResult = decoder.Check(BitResult)
		if len(checkResult) == 0 {
			t.Error("Pass a error data")
		}
		decoder.Iterate()
		BitResult = modem.Demodulate(decoder.GetResult())
		checkResult = decoder.Check(BitResult)
		if len(checkResult) != 0 {
			t.Error("Iterate a right data, but get a wrong answer")
		}
		// Try to solve a simple problem
		BPSKData = modem.Modulate(codeResult)
		BPSKData[0] = -BPSKData[0]
		decoder.LoadData(BPSKData, value)
		BitResult = modem.Demodulate(decoder.GetResult())
		checkResult = decoder.Check(BitResult)
		if len(checkResult) == 0 {
			t.Error("Passed wrong data")
			return
		}
		var itTimes int
		for itTimes = 0; itTimes < 20; itTimes++ {
			decoder.Iterate()
			BitResult = modem.Demodulate(decoder.GetResult())
			checkResult = decoder.Check(BitResult)
			if len(checkResult) == 0 {
				break
			} else {
				t.Logf("Errbits: %v", len(checkResult))
			}
		}
		if itTimes >= 20 {
			t.Error("Cannot solve a simple problem")
			return
		}
		// Try to solve an impossible problem
		BPSKData = modem.Modulate(codeResult)
		for i := 0; i < 500; i++ {
			BPSKData[i] = -BPSKData[i]
		}
		decoder.LoadData(BPSKData, CheckMinSum)
		BitResult = modem.Demodulate(decoder.GetResult())
		checkResult = decoder.Check(BitResult)
		if len(checkResult) == 0 {
			t.Error("Passed wrong data")
			return
		}
		for itTimes = 0; itTimes < 20; itTimes++ {
			decoder.Iterate()
			BitResult = modem.Demodulate(decoder.GetResult())
			checkResult = decoder.Check(BitResult)
			if len(checkResult) == 0 {
				break
			}
		}
		if itTimes < 20 {
			t.Error(fmt.Printf("Solved an impossible problem with %v times?!", itTimes))
			return
		}
		t.Log("Pass")
	}
}
