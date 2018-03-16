package problems

/*
一个球从h高度落下，每次弹起高度都是上次的一半，求十次弹起后的最高高度以及当时球运动距离
*/

func ballfall(h float64, t uint) (curh, toth float64) {
	var i uint = 1

	for i <= t {
		toth += h
		h = h / 2
		i++
	}

	return
}
