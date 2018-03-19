package problems

//有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13...求出这个数列的前20项之和。
func sumfenshu(t int) float64 {
	var up float64 = 2
	var down float64 = 1
	var total float64
	for i := 1; i <= t; i++ {
		total += up / down
		up, down = down, up
		up += down
	}

	return total
}

//或许是要精确些，但分数合并时候对分母乘积计算会出现溢出，结果反而不准确了。
func sumfenshu1(t int) float64 {
	ups := make([]int, t)
	downs := make([]int, t)

	ups[0] = 2
	downs[0] = 1
	var downt float64 = 1
	for i := 1; i < t; i++ {
		ups[i] = ups[i-1] + downs[i-1]
		downs[i] = ups[i-1]
		downt *= float64(downs[i])
	}

	var upt float64
	for i := 0; i < t; i++ {
		upt += float64(ups[i]) * downt / float64(downs[i])
	}

	return float64(upt) / float64(downt)
}
