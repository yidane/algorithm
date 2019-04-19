package problems

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (36.43%)
 * Total Accepted:    287.1K
 * Total Submissions: 787.8K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 *
 *
 * pop() -- Removes the element on top of the stack.
 *
 *
 * top() -- Get the top element.
 *
 *
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 *
 * Example:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 *
 *
 */
type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{}
}

func (minStack *MinStack) Push(x int) {
	minStack.data = append(minStack.data, x)

	if len(minStack.minData) == 0 {
		minStack.minData = append(minStack.minData, x)
		return
	}

	minData := minStack.GetMin()
	if x < minData {
		minStack.minData = append(minStack.minData, x)
	} else {
		minStack.minData = append(minStack.minData, minData)
	}
}

func (minStack *MinStack) Pop() {
	l := len(minStack.minData) - 1
	minStack.data = minStack.data[:l]
	minStack.minData = minStack.minData[0:l]
}

func (minStack *MinStack) Top() int {
	return minStack.data[len(minStack.data)-1]
}

func (minStack *MinStack) GetMin() int {
	return minStack.minData[len(minStack.minData)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
