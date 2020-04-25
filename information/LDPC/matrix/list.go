package matrix

import (
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

type OrderList struct {
	Head *Node
}

func MakeOrderList() OrderList {
	return OrderList{&Node{-1, nil}}
}

func (l *OrderList) Add(v int) {
	cNode := l.Head
	for cNode.Next != nil && cNode.Next.Val <= v {
		cNode = cNode.Next
	}
	if cNode.Val == v {
		return
	}
	nNode := &Node{v, cNode.Next}
	cNode.Next = nNode
}

func (l *OrderList) Delete(v int) {
	cNode := l.Head
	for cNode.Next != nil && cNode.Next.Val < v {
		cNode = cNode.Next
	}
	if cNode.Next != nil && cNode.Next.Val == v {
		cNode.Next = cNode.Next.Next
	}
}

func (l *OrderList) Get(v int) int {
	for cNode := l.Head; cNode != nil && cNode.Val <= v; cNode = cNode.Next {
		if cNode.Val == v {
			return 1
		}
	}
	return 0
}

func (a *OrderList) Compare(b *OrderList) bool {
	ahead := a.Head.Next
	bhead := b.Head.Next
	for ahead != nil && bhead != nil && ahead.Val == bhead.Val {
		ahead = ahead.Next
		bhead = bhead.Next
	}
	return ahead == bhead
}

func (l *OrderList) Sprint() string {
	var sb strings.Builder
	for cnode := l.Head.Next; cnode != nil; cnode = cnode.Next {
		sb.WriteString(strconv.Itoa(cnode.Val))
		sb.WriteString(" ")
	}
	return sb.String()
}

func (l *OrderList) Clone() OrderList {
	result := MakeOrderList()
	tail := result.Head
	for curr := l.Head.Next; curr != nil; curr = curr.Next {
		tail.Next = &Node{curr.Val, nil}
		tail = tail.Next
	}
	return result
}
