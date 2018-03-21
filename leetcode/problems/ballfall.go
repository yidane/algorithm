package problems

import "math"

/*
一个球从h高度落下，每次弹起高度都是上次的一半，求十次弹起后的最高高度以及当时球运动距离
*/

func ballfall(h float64, t uint) (curh, toth float64) {
	var i uint = 1

	curh = h
	for i <= t {
		toth += curh
		curh = curh / 2
		i++
	}

	return
}

/*
构成等比数列，使用等比数列公式计算即可
an=a1*q^n
Sn=a1*(1-q^n)/(1-q) （1!=q）
*/
func ballfall1(h float64, t uint) (curh, toth float64) {
	curh = h * math.Pow(0.5, float64(t))
	toth = h * (1 - math.Pow(0.5, float64(t))) / (1 - 0.5)

	return
}
