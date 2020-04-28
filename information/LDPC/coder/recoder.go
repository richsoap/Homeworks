package coder

import (
	"fmt"
	"math"
)

const (
	ValueSum = iota
	CheckSumProduct
	CheckMinSum
	CheckNormalizedMinSum
	CheckOffsetMinSum
)

func BuildRecorder(recorderType int) interface{} {
	switch recorderType {
	case ValueSum:
		return BuildSumRecorder()
	case CheckMinSum:
		return BuildMinRecorder()
	}
	return nil
}

type Recoder interface {
	Put(index int, value float64)
	Merge()
	Get(index int) float64
	GetMergedResult() float64
	Sprint() string
}

// 和算法(值节点)
type SumRecorder struct {
	values   map[int]float64
	sumValue float64
}

func BuildSumRecorder() *SumRecorder {
	return &SumRecorder{make(map[int]float64), 0}
}

func (r *SumRecorder) Put(index int, value float64) {
	r.values[index] = value
}

func (r *SumRecorder) Merge() {
	r.sumValue = 0
	for _, val := range r.values {
		r.sumValue += val
	}
}

func (r *SumRecorder) Get(index int) float64 {
	res, _ := r.values[index]
	return r.sumValue - res
}

func (r *SumRecorder) GetMergedResult() float64 {
	return r.sumValue
}

func (r *SumRecorder) Sprint() string {
	return fmt.Sprintf("Values in SumRecorder %v", r.values)
}

// 最小和算法(校验节点)
type MinRecorder struct {
	values map[int]float64
	mins   []float64
	sig    float64
}

func BuildMinRecorder() *MinRecorder {
	return &MinRecorder{make(map[int]float64), make([]float64, 0), 1}
}

func (r *MinRecorder) Put(index int, value float64) {
	r.values[index] = value
}

func (r *MinRecorder) Merge() {
	r.mins = []float64{1e7, 1e7}
	r.sig = 1
	for _, value := range r.values {
		temp := math.Abs(value)
		if value < 0 {
			r.sig *= -1
		}
		if temp <= r.mins[0] {
			r.mins[1] = r.mins[0]
			r.mins[0] = temp
		} else if temp < r.mins[1] {
			r.mins[1] = temp
		}
	}
}

func (r *MinRecorder) Get(index int) float64 {
	res, _ := r.values[index]
	sig := r.sig
	if res < 0 {
		sig *= -1
	}
	if math.Abs(res) == r.mins[0] {
		res = r.mins[1]
	} else {
		res = r.mins[0]
	}
	return res * sig
}

func (r *MinRecorder) GetMergedResult() float64 {
	return 0
}

func (r *MinRecorder) Sprint() string {
	return fmt.Sprintf("Values in MinRecorder: %v", r.values)
}
