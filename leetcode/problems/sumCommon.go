package problems

import (
	"fmt"
)

//有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
func sum2() {
	arr := [4]int{1, 2, 3, 4}
	total := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if i == j || i == k || j == k {
					continue
				}
				total++
				fmt.Println(arr[i], arr[j], arr[k])
			}
		}
	}
	fmt.Println("总共组成", total, "个互不相同且无重复数字的三位数")
}

//有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
//1 2，2 2，3 4，4 6，5 10
func sum4(num int) {
	i := 1
	arr := make(map[int]int)
	for {
		if i > num {
			break
		}

		if i < 3 {
			arr[i-1] = 2
		} else {
			arr[i-1] = arr[i-2] + arr[i-3]
		}
		i++
	}

	for i := 0; i < num; i++ {
		fmt.Printf("第%d个月%d\n", i+1, arr[i])
	}
}
