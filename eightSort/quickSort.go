package eightSort

/*
快速排序
	是由东尼·霍尔所发展的一种排序算法。
	在平均状况下，排序 n 个项目要Ο(n log n)次比较。在最坏状况下则需要Ο(n2)次比较，但这种状况并不常见。
	事实上，快速排序通常明显比其他Ο(n log n) 算法更快，因为它的内部循环（inner loop）可以在大部分的架构上很有效率地被实现出来。
	快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。
算法步骤：
	1 从数列中挑出一个元素，称为 “基准”（pivot），
	2 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。
	3 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
	递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。虽然一直递归下去，但是这个算法总会退出，因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。
*/
func quickSort(arr []int) []int {
	quick(arr, 0, len(arr)-1)
	return arr
}

func quick(arr []int, left, right int) {
	//如果left==right,则不需要排序
	if left >= right {
		return
	}

	p, min, max := arr[left], left, right

	//选择最左边元素为基准
	//将比基准小的放在左边
	//比基准大的放在右边
	for min < max {
		for max > min && arr[max] > p {
			max--
		}

		arr[min] = arr[max]

		for max > min && arr[min] <= p {
			min++
		}

		arr[max] = arr[min]
	}

	arr[min] = p

	quick(arr, left, min-1)
	quick(arr, min+1, right)
}

/*
#快速排序 传入列表、开始位置和结束位置
def quick_sort( li , start , end ):
    # 如果start和end碰头了，说明要我排的这个子数列就剩下一个数了，就不用排序了
    if not start < end :
        return

    mid = li[start] #拿出第一个数当作基准数mid
    low = start   #low来标记左侧从基准数始找比mid大的数的位置
    high = end  #high来标记右侧end向左找比mid小的数的位置

    # 我们要进行循环，只要low和high没有碰头就一直进行,当low和high相等说明碰头了
    while low < high :
        #从high开始向左，找到第一个比mid小或者等于mid的数，标记位置,（如果high的数比mid大，我们就左移high）
        # 并且我们要确定找到之前，如果low和high碰头了，也不找了
        while low < high and li[high] > mid :
            high -= 1
        #跳出while后，high所在的下标就是找到的右侧比mid小的数
        #把找到的数放到左侧的空位 low 标记了这个空位
        li[low] = li[high]
        # 从low开始向右，找到第一个比mid大的数，标记位置,（如果low的数小于等于mid，我们就右移low）
        # 并且我们要确定找到之前，如果low和high碰头了，也不找了
        while low < high and li[low] <= mid :
            low += 1
        #跳出while循环后low所在的下标就是左侧比mid大的数所在位置
        # 我们把找到的数放在右侧空位上，high标记了这个空位
        li[high] = li[low]
        #以上我们完成了一次 从右侧找到一个小数移到左侧，从左侧找到一个大数移动到右侧
    #当这个while跳出来之后相当于low和high碰头了，我们把mid所在位置放在这个空位
    li[low] = mid
    #这个时候mid左侧看的数都比mid小，mid右侧的数都比mid大

    #然后我们对mid左侧所有数进行上述的排序
    quick_sort( li , start, low-1 )
    #我们mid右侧所有数进行上述排序
    quick_sort( li , low +1 , end )


#ok我们实践一下
if __name__ == '__main__':
    li = [5,4,3,2,1]
    quick_sort(li , 0 , len(li) -1 )
    print(li)
*/
