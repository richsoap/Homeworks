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

func BuildRecorder(recorderType int, param ...float64) interface{} {
	switch recorderType {
	case ValueSum:
		return BuildSumRecorder()
	case CheckMinSum:
		return BuildMinRecorder()
	case CheckSumProduct:
		return BuildSumProductRecorder()
	case CheckNormalizedMinSum:
		return BuildNorMinRecorder(param[0])
	case CheckOffsetMinSum:
		return BuildOffMinRecorder(param[0])
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
	for i := range r.values {
		r.sumValue += r.values[i]
	}
}

func (r *SumRecorder) Get(index int) float64 {
	if res, ok := r.values[index]; ok {
		return r.sumValue - res
	} else {
		return 0
	}
}

func (r *SumRecorder) GetMergedResult() float64 {
	return r.sumValue
}

func (r *SumRecorder) Sprint() string {
	return fmt.Sprintf("Values in SumRecorder %v", r.values)
}

// 和积算法（校验节点）
type SumProductRecorder struct {
	values  map[int]float64
	product float64
}

func BuildSumProductRecorder() *SumProductRecorder {
	return &SumProductRecorder{make(map[int]float64), 0}
}

func (r *SumProductRecorder) Put(index int, value float64) {
	r.values[index] = value
}

func (r *SumProductRecorder) Merge() {
	r.product = 1
	for i := range r.values {
		r.product *= math.Tanh(r.values[i] / 2)
	}
}

func (r *SumProductRecorder) Get(index int) float64 {
	if val, ok := r.values[index]; ok {
		res := 2 * math.Atanh(r.product/math.Tanh(val/2))
		if math.IsInf(res, 1) {
			return 100
		} else if math.IsInf(res, 0) {
			return -100
		} else {
			return res
		}
	} else {
		return 0
	}
}

func (r *SumProductRecorder) GetMergedResult() float64 {
	return 0
}

func (r *SumProductRecorder) Sprint() string {
	return fmt.Sprintf("Values in SumProductRecorder: %v", r.values)
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

func (r *MinRecorder) Merge() {}

func (r *MinRecorder) Get(index int) float64 {
	result := 1e7
	sig := 1.0
	for i, val := range r.values {
		if i == index {
			continue
		}
		if result > math.Abs(val) {
			result = math.Abs(val)
		}
		if val < 0 {
			sig *= -1
		}
	}
	return result * sig
}

func (r *MinRecorder) GetMergedResult() float64 {
	return 0
}

func (r *MinRecorder) Sprint() string {
	return fmt.Sprintf("Values in MinRecorder: %v", r.values)
}

// 归一化最小和算法(校验节点)
type NorMinRecorder struct {
	values map[int]float64
	mins   []float64
	sig    float64
	a      float64
}

func BuildNorMinRecorder(a float64) *NorMinRecorder {
	return &NorMinRecorder{make(map[int]float64), make([]float64, 0), 1, a}
}

func (r *NorMinRecorder) Put(index int, value float64) {
	r.values[index] = value
}

func (r *NorMinRecorder) Merge() {}

func (r *NorMinRecorder) Get(index int) float64 {
	result := 1e7
	sig := 1.0
	for i, val := range r.values {
		if i == index {
			continue
		}
		if result > math.Abs(val) {
			result = math.Abs(val)
		}
		if val < 0 {
			sig *= -1
		}
	}
	return result * sig * r.a
}

func (r *NorMinRecorder) GetMergedResult() float64 {
	return 0
}

func (r *NorMinRecorder) Sprint() string {
	return fmt.Sprintf("Values in MinRecorder: %v", r.values)
}

// 偏置最小和算法(校验节点)
type OffMinRecorder struct {
	values map[int]float64
	b      float64
}

func BuildOffMinRecorder(b float64) *OffMinRecorder {
	return &OffMinRecorder{make(map[int]float64), b}
}

func (r *OffMinRecorder) Put(index int, value float64) {
	r.values[index] = value
}

func (r *OffMinRecorder) Merge() {}

func (r *OffMinRecorder) Get(index int) float64 {
	result := 1e7
	sig := 1.0
	for i, val := range r.values {
		if i == index {
			continue
		}
		if result > math.Abs(val) {
			result = math.Abs(val)
		}
		if val < 0 {
			sig *= -1
		}
	}
	return math.Max(0, result-r.b) * sig
}

func (r *OffMinRecorder) GetMergedResult() float64 {
	return 0
}

func (r *OffMinRecorder) Sprint() string {
	return fmt.Sprintf("Values in MinRecorder: %v", r.values)
}
