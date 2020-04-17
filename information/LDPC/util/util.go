package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func FloatSliceCompare(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func IntSliceCompare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func LoadHMatrix(path string) ([][]int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	sizeStr, err := reader.ReadString('\n')
	matrixSize := strings.Fields(sizeStr)
	if len(matrixSize) != 3 {
		return nil, 0, errors.New(fmt.Sprintf("size info error, get %v", len(matrixSize)))
	}
	h, err := strconv.Atoi(matrixSize[0])
	if err != nil {
		return nil, 0, errors.New(fmt.Sprintf("h conv errors %v", matrixSize[0]))
	}
	w, err := strconv.Atoi(matrixSize[1])
	if err != nil {
		return nil, 0, errors.New(fmt.Sprintf("w conv errors %v", matrixSize[1]))
	}
	z, err := strconv.Atoi(matrixSize[2])
	if err != nil {
		return nil, 0, errors.New(fmt.Sprintf("z conv errors %v", matrixSize[2]))
	}
	result := make([][]int, h, w)
	index := 0
	for {
		valueStr, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, 0, errors.New(fmt.Sprintf("line %v: read error", index))
		}
		if index >= h {
			return nil, 0, errors.New("too many lines")
		}
		ValueStrs := strings.Fields(valueStr)
		if len(ValueStrs) != w {
			return nil, 0, errors.New(fmt.Sprintf("line %v: not enough value need %v get %v", index, w, len(ValueStrs)))
		}
		result[index] = make([]int, w, w)
		for i := 0; i < w; i++ {
			result[index][i], err = strconv.Atoi(ValueStrs[i])
			if err != nil {
				return nil, 0, errors.New(fmt.Sprintf("line %v %v: conv error %v", index, i, ValueStrs[i]))
			}
		}
		index++
	}
	if index != h {
		return nil, 0, errors.New(fmt.Sprintf("not enough lines need %v get %v", h, index))
	}
	return result, z, nil
}

func IntSlicePrint(value []int) {
	for _, v := range value {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
