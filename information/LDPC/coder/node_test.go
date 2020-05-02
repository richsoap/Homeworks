package coder

import (
	"fmt"
	"testing"

	"github.com/richsoap/ldpc/util"
)

func LogNodesDetails(t *testing.T, nodes *[]Node) {
	for i := 0; i < len(*nodes); i++ {
		t.Error(fmt.Sprintf("Value Node %v: self val: %v, merged val: %v, %v", i, (*nodes)[i].val, (*nodes)[i].rec.GetMergedResult(), (*nodes)[i].rec.Sprint()))
	}
}

func TestNode(t *testing.T) {
	valueNodes := make([]Node, 0)
	sig := 1.0
	for i := 0; i < 5; i++ {
		valueNodes = append(valueNodes, BuildNode(i, sig*float64(i), BuildSumRecorder()))
		sig *= -1
	}
	checkNodes := make([]Node, 0)
	for i := 0; i < 5; i++ {
		checkNodes = append(checkNodes, BuildNode(i, 0, BuildMinRecorder()))
	}
	for i := 0; i < 5; i++ {
		valueNodes[i].AddLink(i)
		valueNodes[i].AddLink((i + 4) % 5)
		checkNodes[i].AddLink(i)
		checkNodes[i].AddLink((i + 1) % 5)
	}
	for i := 0; i < 5; i++ {
		checkNodes[i].Update(&valueNodes, nil)
	}
	for i := 0; i < 5; i++ {
		valueNodes[i].Update(&checkNodes, nil)
	}
	result := make([]float64, 5, 5)
	target := []float64{3, 1, -2, 3, 1}
	for i := 0; i < 5; i++ {
		result[i] = valueNodes[i].GetResult()
	}
	if !util.FloatSliceCompare(result, target) {
		t.Error(fmt.Sprintf("First Float Compare Error: wanted %v, get %v", target, result))
		t.Error("Value Nodes")
		LogNodesDetails(t, &valueNodes)
		t.Error("Check Nodes")
		LogNodesDetails(t, &checkNodes)
	}
	for i := 0; i < 5; i++ {
		checkNodes[i].Update(&valueNodes, nil)
	}
	for i := 0; i < 5; i++ {
		valueNodes[i].Update(&checkNodes, nil)
	}
	target = []float64{2, 2, 2, 2, 2}
	for i := 0; i < 5; i++ {
		result[i] = valueNodes[i].GetResult()
	}
	if !util.FloatSliceCompare(result, target) {
		t.Error(fmt.Sprintf("Second Float Compare Error: wanted %v, get %v", target, result))
		t.Error("Value Nodes")
		LogNodesDetails(t, &valueNodes)
		t.Error("Check Nodes")
		LogNodesDetails(t, &checkNodes)
	}
}
