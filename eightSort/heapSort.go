package eightSort

/*

https://blog.csdn.net/u012279631/article/details/80712807

堆排序（Heapsort）
	是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
	堆排序的平均时间复杂度为Ο(nlogn) 。
算法步骤：
	1）创建一个堆H[0..n-1]
	2）把堆首（最大值）和堆尾互换
	3）把堆的尺寸缩小1，并调用shift_down(0),目的是把新的数组顶端数据调整到相应位置
	4） 重复步骤2，直到堆的尺寸为1
*/

func heapSort(arr []int) []int {
	var n = len(arr) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, arr)
	}

	for end := n; end >= 0; end-- {
		if arr[0] < arr[end] {
			arr[0], arr[end] = arr[end], arr[0]
			minHeap(0, end-1, arr)
		}
	}

	return arr
}

func minHeap(root int, end int, c []int) {
	for {
		//root的子节点位于 2*root+1和2*root+2
		var child = 2*root + 1
		//判断是否存在child节点
		if child > end {
			break
		}

		//判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}

		//构建最小堆
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}
