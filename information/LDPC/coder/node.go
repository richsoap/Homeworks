package coder

import "sync"

type Node struct {
	index int
	val   float64
	rec   Recoder
	links []int
}

func BuildNode(index int, val float64, rec Recoder) Node {
	result := Node{index, val, rec, make([]int, 0)}
	return result
}

func (n *Node) Get(index int) float64 {
	return n.val + n.rec.Get(index)
}

func (n *Node) Update(nodes *[]Node, wg *sync.WaitGroup) {
	for _, i := range n.links {
		n.rec.Put(i, (*nodes)[i].Get(n.index))
	}
	n.rec.Merge()
	if wg != nil {
		wg.Done()
	}
}

func (n *Node) AddLink(index int) {
	n.links = append(n.links, index)
}

func (n *Node) GetResult() float64 {
	return n.val + n.rec.GetMergedResult()
}
