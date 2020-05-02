package matrix

import (
	"fmt"
	"testing"

	"github.com/richsoap/ldpc/source"
	"github.com/richsoap/ldpc/util"
)

func checkValue(t *testing.T, b *BinaryMatrix, r, c, v int, ok bool) {
	if rv, rok := b.Get(r, c); rok != ok {
		t.Error(fmt.Sprintf("State error %v %v, get %v", r, c, rok))
	} else if rv != v {
		t.Error(fmt.Sprintf("Result error %v %v, wanted %v, get %v", r, c, v, rv))
	}
}

func TestMatrix(t *testing.T) {
	m := MakeBinaryMatrix(3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checkValue(t, &m, i, j, 0, true)
		}
	}
	for i := 0; i < 3; i++ {
		if ok := m.Set(i, 0, 1); !ok {
			t.Error(fmt.Sprintf("Set error %v %v", i, 0))
		}
	}
	for i := 4; i < 6; i++ {
		if ok := m.Set(i, 0, 1); ok {
			t.Error(fmt.Sprintf("Set error %v %v", i, 0))
		}
	}
	for i := 0; i < 3; i++ {
		checkValue(t, &m, i, 0, 1, true)
	}
	for i := 3; i < 9; i++ {
		checkValue(t, &m, i, 0, 0, false)
	}
	for i := 0; i < 3; i++ {
		if ok := m.Set(i, 0, 0); !ok {
			t.Error(fmt.Sprintf("Set error %v %v", i, 0))
		}
	}
	for i := 0; i < 3; i++ {
		checkValue(t, &m, i, 0, 0, true)
	}
}

func checkMul(t *testing.T, a, b, c *BinaryMatrix) {
	if res, err := a.Mul(b); err != nil {
		t.Error(err)
	} else if !res.Compare(c) {
		t.Error("Mul result error")
		t.Error(a.Sprint())
		t.Error("******************")
		t.Error(b.Sprint())
		t.Error("====================")
		t.Error(res.Sprint())
		t.Error("-----------------")
		t.Error(c.Sprint())
	}
}

func TestMul(t *testing.T) {
	av := [][]int{{0, 1, 0, 1, 1}}
	bv := [][]int{{1, 0, 1, 0, 1}, {1, 1, 0, 0, 1}, {1, 0, 0, 1, 1}, {1, 0, 1, 1, 1}}
	cv := [][]int{{1}, {0}, {0}, {0}}
	a := MakeBinaryMatrixWithData(av)
	b := MakeBinaryMatrixWithData(bv)
	c := MakeBinaryMatrixWithData(cv)
	a = a.Trans()
	checkMul(t, &b, &a, &c)
	av = [][]int{{0, 0, 0, 1, 1, 0, 1, 0}}
	bv = [][]int{{1, 1, 0, 1, 1, 0, 1, 0}}
	cv = [][]int{{1}}
	a = MakeBinaryMatrixWithData(av)
	b = MakeBinaryMatrixWithData(bv)
	c = MakeBinaryMatrixWithData(cv)
	b = b.Trans()
	checkMul(t, &a, &b, &c)
}

func TestTrans(t *testing.T) {
	av := [][]int{{1, 0, 1}}
	bv := [][]int{{1}, {0}, {1}}
	a := MakeBinaryMatrixWithData(av)
	b := MakeBinaryMatrixWithData(bv)
	aT := a.Trans()
	if !aT.Compare(&b) {
		t.Error("Trans Error")
		t.Error("Wanted:")
		t.Error(b.Sprint())
		t.Error("Get:")
		t.Error(aT.Sprint())
	}
}

func TestGetLine(t *testing.T) {
	av := make([][]int, 3)
	for i := range av {
		av[i] = source.Create(10000)
	}
	a := MakeBinaryMatrixWithData(av)
	a = a.Trans()
	a = a.Trans()
	for i := range av {
		ar, err := a.GetLine(i)
		if err != nil {
			t.Error(fmt.Sprintf("line %v Error: %v", i, err))
		}
		if !util.IntSliceCompare(av[i], ar) {
			t.Error(fmt.Sprintf("line compare %v error", i))
		}
	}
}
