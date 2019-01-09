package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one *ListNode
	two *ListNode
}

type answer struct {
	one *ListNode
}

type question struct {
	p param
	a answer
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: param{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: answer{
				one: makeListNode([]int{7, 0, 8}),
			},
		},
		question{
			p: param{
				one: makeListNode([]int{1, 2, 3}),
				two: makeListNode([]int{1, 2, 3}),
			},
			a: answer{
				one: makeListNode([]int{2, 4, 6}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addTwoNumbers(p.one, p.two), "输入:%v", p)
	}
}
