package matrix

import (
	"errors"
	"fmt"
	"strings"
)

type BinaryMatrix struct {
	W   int
	H   int
	Row []OrderList
	Col []OrderList
}

func MakeBinaryMatrix(h, w int) BinaryMatrix {
	res := BinaryMatrix{}
	res.W = w
	res.H = h
	res.Row = make([]OrderList, h, h)
	res.Col = make([]OrderList, w, w)
	for i := 0; i < h; i++ {
		res.Row[i] = MakeOrderList()
	}
	for i := 0; i < w; i++ {
		res.Col[i] = MakeOrderList()
	}
	return res
}

func MakeBinaryMatrixWithData(data [][]int) BinaryMatrix {
	res := MakeBinaryMatrix(len(data), len(data[0]))
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] != 0 {
				res.Set(i, j, 1)
			}
		}
	}
	return res
}

func (m *BinaryMatrix) Set(r, c, val int) bool {
	if c >= m.W || r >= m.H {
		return false
	}
	if val == 0 {
		m.Row[r].Delete(c)
		m.Col[c].Delete(r)
	} else {
		m.Row[r].Add(c)
		m.Col[c].Add(r)
	}
	return true
}

func (m *BinaryMatrix) Get(r, c int) (int, bool) {
	if r >= m.H || c >= m.W {
		return 0, false
	}
	return m.Row[r].Get(c), true
}

func (a *BinaryMatrix) Compare(b *BinaryMatrix) bool {
	if a.W != b.W || a.H != b.H {
		return false
	}
	result := true
	for i := 0; i < len(a.Row) && result; i++ {
		result = a.Row[i].Compare(&b.Row[i])
	}
	for i := 0; i < len(a.Col) && result; i++ {
		result = a.Col[i].Compare(&b.Col[i])
	}
	return result
}

type mulResult struct {
	r, c, v int
}

func (m *BinaryMatrix) Mul(nm *BinaryMatrix) (BinaryMatrix, error) {
	if m.W != nm.H {
		return BinaryMatrix{}, errors.New(fmt.Sprintf("Size not match %v*%v %v*%v", m.H, m.W, nm.H, nm.W))
	}
	result := MakeBinaryMatrix(m.H, nm.W)
	resChan := make(chan mulResult)
	cal := func(a, b *BinaryMatrix, r, c int) {
		record := make(map[int]int)
		count := 0
		for node := a.Row[r].Head.Next; node != nil; node = node.Next {
			record[node.Val] = 1
		}
		for node := b.Col[c].Head.Next; node != nil; node = node.Next {
			if _, ok := record[node.Val]; ok {
				count += 1
			}
		}
		resChan <- mulResult{r, c, count % 2}
	}
	for i := 0; i < result.H; i++ {
		for j := 0; j < result.W; j++ {
			go cal(m, nm, i, j)
		}
	}

	for i := 0; i < result.H*result.W; i++ {
		res := <-resChan
		if res.v == 1 {
			result.Set(res.r, res.c, res.v)
		}
	}
	close(resChan)
	return result, nil
}

func (m *BinaryMatrix) GetLine(index int) ([]int, error) {
	if index >= m.H {
		return nil, errors.New(fmt.Sprintf("Get line out of range %v", index))
	}
	result := make([]int, m.W, m.W)
	for i := 0; i < m.W; i++ {
		result[i] = 0
	}
	for head := m.Row[index].Head.Next; head != nil; head = head.Next {
		result[head.Val] = 1
	}
	return result, nil
}

func (m *BinaryMatrix) Sprint() string {
	var sb strings.Builder
	sb.WriteString("\n")
	for i := 0; i < len(m.Row); i++ {
		sb.WriteString(fmt.Sprintf("%v: ", i))
		sb.WriteString(m.Row[i].Sprint())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m *BinaryMatrix) Trans() BinaryMatrix {
	result := MakeBinaryMatrix(m.W, m.H)
	for i := 0; i < m.W; i++ {
		result.Row[i] = m.Col[i].Clone()
	}
	for i := 0; i < m.H; i++ {
		result.Col[i] = m.Row[i].Clone()
	}
	return result
}
