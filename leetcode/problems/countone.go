// 有一个整数n,写一个函数f(n),返回0到n之间出现的"1"的个数。
// 比如f(13)=6,现在f(1)=1,问下一个最大的f(n)=n的n是什么？

package problems

import (
	"fmt"
	"time"
)

var cache map[int]int = make(map[int]int)
var totalCache map[int]int = make(map[int]int)

//Fn 返回0到max之间出现的“1”的个数
func Fn(max int) int {
	if max <= 0 {
		return 0
	}
	if max == 1 {
		totalCache[1] = 1
		return 1
	}
	var result = CountOne(max)
	cache[max] = CountOne(max)
	result += totalCache[max-1]

	totalCache[max] = result
	return result
}

/*
1:1
2:1
3:1
4:1
5:1
6:1
7:1
8:1
9:1
10:2
11:4
12:4
13:5
14:6
15:7
19:11
20:11
*/

func CountOne(num int) int {
	if _, ok := cache[num]; ok {
		return cache[num]
	}
	if num < 1 {
		return 0
	}
	if num == 1 || num == 10 {
		return 1
	}
	if num < 10 && num > 1 {
		return 1
	}
	var result int
	m := num % 10
	if m == 1 {
		result++
	}
	result = result + CountOne((num-num%10)/10)

	return result
}

func CountOne0(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		t := i
		for t >= 10 {
			j := t % 10
			if j == 1 {
				total += 1
			}
			t = (t - j) / 10
		}

		if t == 1 {
			total += 1
		}
	}

	return total
}

func CountOne1(num int) int {
	hCount := 0

	n := num % 10
	m := (num - n) / 10
	l := m
	for m > 0 {
		if m%10 == 1 {
			hCount++
		}
		m = (m - m%10) / 10
	}

	total := 0
	for i := 0; i <= n; i++ {
		if i == 1 {
			total++
		}
		if hCount > 0 {
			total += hCount
		}
	}

	if l > 10 {
		return total + CountOne1(l)
	}
	return total
}

//CountOneExec 输出200000以内所有F(n)=n的数
func CountOneExec() {
	beginTime := time.Now()
	fmt.Println("begin")
	for i := 0; i < 200000; i++ {
		if i == Fn(i) {
			fmt.Println(i)
		}
	}
	fmt.Println("time:", time.Now().Sub(beginTime).String())
	fmt.Println("finish")
}
