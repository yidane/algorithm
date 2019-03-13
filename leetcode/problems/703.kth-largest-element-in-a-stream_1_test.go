package problems

import (
	"testing"
)

func TestNewKthLargest(t *testing.T) {
	/*
		int k = 3;
		int[] arr = [4,5,8,2];
		KthLargest kthLargest = new KthLargest(3, arr);
		kthLargest.add(3);   // returns 4
		kthLargest.add(5);   // returns 5
		kthLargest.add(10);  // returns 5
		kthLargest.add(9);   // returns 8
		kthLargest.add(4);   // returns 8
	*/

	arr := []int{4, 5, 8, 2}
	k := 3

	kl := NewKthLargest(k, arr)

	testCases := []struct{ k, v int }{
		{k: 3, v: 4},
		{k: 5, v: 5},
		{k: 10, v: 5},
		{k: 9, v: 8},
		{k: 4, v: 8},
	}

	for _, tt := range testCases {
		result := kl.Add(tt.k)
		if result != tt.v {
			t.Fatalf("k(%v) != v(%v)", result, tt.v)
		}
	}
}

func TestNewKthLargest1(t *testing.T) {
	/*
		["KthLargest","add","add","add","add","add"]
		[[1,[]],[-3],[-2],[-4],[0],[4]]
	*/

	var arr []int
	k := 1

	kl := NewKthLargest(k, arr)

	testCases := []struct{ k, v int }{
		{k: -3, v: -3},
		{k: -2, v: -2},
		{k: -4, v: -2},
		{k: 0, v: 0},
		{k: 4, v: 4},
	}

	for _, tt := range testCases {
		result := kl.Add(tt.k)
		if result != tt.v {
			t.Fatalf("k(%v) != v(%v)", result, tt.v)
		}
	}
}

func TestNewKthLargest2(t *testing.T) {
	/*
		["KthLargest","add","add","add","add","add"]
		[[2,[0]],[-1],[1],[-2],[-4],[3]]
	*/
	arr := []int{0}
	k := 2

	kl := NewKthLargest(k, arr)

	testCases := []struct{ k, v int }{
		{k: -1, v: -1},
		{k: 1, v: 0},
		{k: -2, v: 0},
		{k: -4, v: 0},
		{k: 3, v: 1},
	}

	for _, tt := range testCases {
		result := kl.Add(tt.k)
		if result != tt.v {
			t.Fatalf("k(%v) != v(%v)", result, tt.v)
		}
	}
}
