package main

import (
	"fmt"
	"log"
	"os"

	"github.com/richsoap/ldpc/channel"
	"github.com/richsoap/ldpc/coder"
	"github.com/richsoap/ldpc/modem"
	"github.com/richsoap/ldpc/source"
	"github.com/richsoap/ldpc/util"
)

func main() {
	//a, _ := GetBestAlpha()
	b, _ := GetBestBeta()
	//RunTest("SP", coder.CheckSumProduct)
	//RunTest("MS", coder.CheckMinSum)
	//RunTest("NMS", coder.CheckNormalizedMinSum, a)
	RunTest("OMS", coder.CheckOffsetMinSum, b)
}

func RunTest(name string, rec int, param ...float64) {
	HMatrix, z, err := util.LoadHMatrix("./doc/H_block.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	sigLen := len(HMatrix)
	encoder := coder.MakeEncoder(HMatrix, len(HMatrix), len(HMatrix[0]), z)
	decoder := coder.MakeDecoder(HMatrix, len(HMatrix), len(HMatrix[0]))
	noise := []float64{-1, -0.5, 0, 0.5, 1, 1.5, 2}
	errframe := []int{50, 50, 50, 50, 50, 3, 3}
	fileObj, _ := os.Create(fmt.Sprintf("./data/%v.csv", name))
	fmt.Fprintln(fileObj, "frame,errframe,errbits,type,snr")
	for i := range noise {
		log.Printf("%v: %v", name, noise[i])
		errframecount := 0
		errbitscount := 0
		framecount := 0
		awgnChannel := channel.BuildChannel(noise[i])
		for errframecount < errframe[i] {
			framecount++
			if framecount%100 == 0 {
				print(framecount)
				print(errframecount)
			}
			sourceData := source.Create(sigLen)
			codeResult, err := encoder.Encode(sourceData)
			if err != nil {
				log.Fatalf("Encode Error %v", err)
				return
			}
			BPSKData := modem.Modulate(codeResult)
			BPSKData = awgnChannel.AWGN(BPSKData)
			sigmaSq := channel.FromDB(noise[i])
			for j := range BPSKData {
				BPSKData[j] = BPSKData[j] * 2 * sigmaSq
			}
			decoder.LoadData(BPSKData, rec, param...)
			for j := 0; j < 30; j++ {
				decoder.Iterate()
				BitResult := modem.Demodulate(decoder.GetResult())
				checkResult := decoder.Check(BitResult)
				if len(checkResult) == 0 {
					break
				}
			}
			BitResult := modem.Demodulate(decoder.GetResult())
			errbits := 0
			for j := range sourceData {
				if BitResult[j+1008] != sourceData[j] {
					errbits++
				}
			}
			if errbits != 0 {
				errframecount++
			}
			errbitscount += errbits
		}
		fmt.Fprintf(fileObj, "%v,%v,%v,%v,%v\n", framecount, errframecount, errbitscount, name, noise[i])
	}

}

func GetBestAlpha() (float64, error) {
	HMatrix, z, err := util.LoadHMatrix("./doc/H_block.txt")
	if err != nil {
		return 0, err
	}
	sigLen := len(HMatrix)
	noise := 1.0
	bestebr := 1e7
	bestalpha := -1.0
	encoder := coder.MakeEncoder(HMatrix, len(HMatrix), len(HMatrix[0]), z)
	decoder := coder.MakeDecoder(HMatrix, len(HMatrix), len(HMatrix[0]))
	awgnChannel := channel.BuildChannel(noise)

	fileObj, err := os.Create("./data/alpha.csv")
	if err != nil {
		return 0, err
	}
	fmt.Fprintln(fileObj, "alpha,ebr")
	for i := 0.0; i <= 1.0; i += 0.1 {
		errbits := 0
		log.Printf("alpha: %v", i)
		for j := 0; j < 100; j++ {
			sourceData := source.Create(sigLen)
			codeResult, _ := encoder.Encode(sourceData)
			BPSKData := modem.Modulate(codeResult)
			BPSKData = awgnChannel.AWGN(BPSKData)
			sigmaSq := channel.FromDB(noise)
			for j := range BPSKData {
				BPSKData[j] = BPSKData[j] * 2 * sigmaSq
			}
			decoder.LoadData(BPSKData, coder.CheckNormalizedMinSum, i)
			for j := 0; j < 30; j++ {
				decoder.Iterate()
				BitResult := modem.Demodulate(decoder.GetResult())
				checkResult := decoder.Check(BitResult)
				if len(checkResult) == 0 {
					break
				}
			}
			BitResult := modem.Demodulate(decoder.GetResult())
			for j := range sourceData {
				if BitResult[j+1008] != sourceData[j] {
					errbits++
				}
			}
		}
		ebr := float64(errbits) / 1008000.0
		fmt.Fprintf(fileObj, "%v,%v\n", i, ebr)
		if ebr < bestebr {
			bestebr = ebr
			bestalpha = i
		}
	}
	return bestalpha, nil
}
func GetBestBeta() (float64, error) {
	HMatrix, z, err := util.LoadHMatrix("./doc/H_block.txt")
	if err != nil {
		return 0, err
	}
	sigLen := len(HMatrix)
	noise := 1.5
	bestebr := 1e7
	bestalpha := -1.0
	encoder := coder.MakeEncoder(HMatrix, len(HMatrix), len(HMatrix[0]), z)
	decoder := coder.MakeDecoder(HMatrix, len(HMatrix), len(HMatrix[0]))
	awgnChannel := channel.BuildChannel(noise)

	fileObj, err := os.Create("./data/beta.csv")
	if err != nil {
		return 0, err
	}
	fmt.Fprintln(fileObj, "beta, ebr")
	for i := 0.0; i <= 1.0; i += 0.1 {
		errbits := 0
		log.Printf("beta: %v", i)
		for j := 0; j < 100; j++ {
			sourceData := source.Create(sigLen)
			codeResult, _ := encoder.Encode(sourceData)
			BPSKData := modem.Modulate(codeResult)
			BPSKData = awgnChannel.AWGN(BPSKData)
			sigmaSq := channel.FromDB(noise)
			for j := range BPSKData {
				BPSKData[j] = BPSKData[j] * 2 * sigmaSq
			}
			decoder.LoadData(BPSKData, coder.CheckNormalizedMinSum, i)
			for j := 0; j < 30; j++ {
				decoder.Iterate()
				BitResult := modem.Demodulate(decoder.GetResult())
				checkResult := decoder.Check(BitResult)
				if len(checkResult) == 0 {
					break
				}
			}
			BitResult := modem.Demodulate(decoder.GetResult())
			for j := range sourceData {
				if BitResult[j+1008] != sourceData[j] {
					errbits++
				}
			}
		}
		ebr := float64(errbits) / 1008000.0
		fmt.Fprintf(fileObj, "%v, %v\n", i, ebr)
		if ebr < bestebr {
			bestebr = ebr
			bestalpha = i
		}
	}
	return bestalpha, nil
}
