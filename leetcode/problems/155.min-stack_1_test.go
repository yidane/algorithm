package problems

import (
	"testing"
)

func TestNewMinStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if minStack.GetMin() != -3 {
		t.Fatal()
	}

	minStack.Pop()
	if minStack.Top() != 0 {
		t.Fatal()
	}

	if minStack.GetMin() != -2 {
		t.Fatal()
	}

	//["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

	ms := NewMinStack()
	ms.Push(2147483646)
	ms.Push(2147483646)
	ms.Push(2147483646)
	ms.Top()
	ms.Pop()
	ms.GetMin()
	ms.Top()
	ms.GetMin()
	ms.Pop()
	ms.Push(2147483647)
	ms.Top()
	ms.GetMin()
	ms.Push(-2147483648)
	ms.Top()
	ms.GetMin()
	ms.Pop()
	ms.GetMin()
}
