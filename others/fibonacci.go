package others

//https://www.itcodemonkey.com/article/12793.html

/*
斐波那契数列求解
*/

//递归求解，计算速度很慢
func fibonacci(size int) int {
	if size <= 1 {
		return 1
	}

	return fibonacci(size-1) + fibonacci(size-2)
}

//迭代算法
func fibonacci1(size int) int {
	p0, p1 := 0, 1
	n := 1
	for n < size {
		p0, p1 = p1, p0
		p1 = p0 + p1
		n++
	}

	return p0 + p1
}
