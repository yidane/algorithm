package others

import (
	"fmt"
	"github.com/yidane/fraction"
	"strings"
)

/*
https://leetcode.com/problems/24-game/description/
You have 4 cards each containing a number from 1 to 9.
You need to judge whether they could operated through *, /, +, -, (, ) to get the value of 24.

Example 1:
Input: [4, 1, 8, 7]
Output: True
Explanation: (8-4) * (7-1) = 24
Example 2:
Input: [1, 2, 1, 2]
Output: False
Note:
The division operator / represents real division, not integer division. For example, 4 / (1 - 2/3) = 12.
Every operation done is between two numbers. In particular, we cannot use - as a unary operator.
For example,with [1, 1, 1, 1] as input,the expression -1 - 1 - 1 - 1 is not allowed.
You cannot concatenate numbers together. For example, if the input is [1, 2, 1, 2], we cannot write this as 12 + 12.
*/

/*
+
-
×
	2*12
	3*8
	4*6
/

*/

//todo 尚未完全实现
func judgePoint24(nums []int) bool {
	r, s, t, u, target := newFraction(nums[0]), newFraction(nums[1]), newFraction(nums[2]), newFraction(nums[3]), newFraction(24)

	var exps []string
	for _, n := range newTarget(r, target) {
		newTarget := n
		es := judgePoint3(s, t, u, newTarget.v)
		exps = append(exps, packExp(es, newTarget)...)
	}
	for _, n := range newTarget(s, target) {
		newTarget := n
		es := judgePoint3(r, t, u, newTarget.v)
		exps = append(exps, packExp(es, newTarget)...)

	}
	for _, n := range newTarget(t, target) {
		newTarget := n
		es := judgePoint3(r, s, u, newTarget.v)
		exps = append(exps, packExp(es, newTarget)...)
	}
	for _, n := range newTarget(u, target) {
		newTarget := n
		es := judgePoint3(r, s, t, newTarget.v)
		exps = append(exps, packExp(es, newTarget)...)
	}

	return len(exps) > 0
}

//判断3个数能否满足等于目标值
func judgePoint3(r, s, t, target fraction.Fraction) (exps []string) {
	for _, n := range newTarget(r, target) {
		newTarget := n
		es := judgePoint2(s, t, newTarget.v)
		exps = append(exps, packExp(es, newKV(newTarget.k, r))...)
	}

	for _, n := range newTarget(s, target) {
		newTarget := n
		es := judgePoint2(r, t, newTarget.v)
		exps = append(exps, packExp(es, newKV(newTarget.k, s))...)
	}

	for _, n := range newTarget(t, target) {
		newTarget := n
		es := judgePoint2(r, s, newTarget.v)
		exps = append(exps, packExp(es, newKV(newTarget.k, t))...)
	}

	return
}

//判断2个数能否满足等于目标值
func judgePoint2(r, s, target fraction.Fraction) (exps []string) {
	var p fraction.Fraction
	if p = s.Add(r); p.Equal(target) {
		exps = append(exps, fmt.Sprintf("%v+%v", s, r))
	}

	if p = s.Subtract(r); p.Equal(target) {
		exps = append(exps, fmt.Sprintf("%v-%v", s, r))
	}

	if p = r.Subtract(s); p.Equal(target) {
		exps = append(exps, fmt.Sprintf("%v-%v", r, s))
	}

	if p = r.Multiply(s); p.Equal(target) {
		exps = append(exps, fmt.Sprintf("%v*%v", s, r))
	}

	if p, _ = s.Divide(r); p.Equal(target) {
		exps = append(exps, fmt.Sprintf("%v÷%v", s, r))
	}

	if p, _ = r.Divide(s); p.Equal(target) {
		exps = append(exps, fmt.Sprintf("%v÷%v", r, s))
	}

	return
}

func newFraction(i int) fraction.Fraction {
	f, _ := fraction.New(int64(i), 1)
	return f
}

type kv struct {
	k rune
	v fraction.Fraction
}

func packExp(exps []string, op kv) []string {
	if len(exps) == 0 {
		return exps
	}

	var result []string
	for i := 0; i < len(exps); i++ {
		exp := exps[i]
		switch op.k {
		case '+':
			result = append(result, exp)
		case '-':
			result = append(result, exp)
		case '×':
			if strings.IndexRune(exp, '+') > 0 || strings.IndexRune(exp, '-') > 0 {
				exp = "(" + exp + ")"
			}

			result = append(result, exp+string(op.k)+op.v.String())
			result = append(result, op.v.String()+string(op.k)+exp)
		case '÷':
			if strings.IndexRune(exp, '+') > 0 || strings.IndexRune(exp, '-') > 0 {
				exp = "(" + exp + ")" + string(op.k) + op.v.String()
			}

			result = append(result, exp)
		}
	}

	return result
}

func newKV(k rune, v fraction.Fraction) kv {
	return kv{k: k, v: v}
}

func newTarget(r, target fraction.Fraction) []kv {
	result := make([]kv, 0, 4)

	result = append(result, newKV('-', target.Add(r)))      //加
	result = append(result, newKV('+', target.Subtract(r))) //减法
	result = append(result, newKV('÷', target.Multiply(r))) //乘法

	df, _ := target.Divide(r)
	result = append(result, newKV('×', df)) //除法

	return result
}
