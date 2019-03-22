package bytedance

/*
1 1 0 1 1
1 0 1 1 1
0 0 0 1 1
1 1 0 0 1
1 1 1 0 1
*/

func findGroup(arr [][]int) (count int, maxGroup int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][j] == 1 {

			}
		}
	}

	return
}
