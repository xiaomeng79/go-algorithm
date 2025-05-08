package quick

//快速排序:快速排序应该算是在冒泡排序基础上的递归分治法
//主要思想：之所以快，跳跃式的交换消除逆序，

//算法步骤:
//从数列中挑出一个元素，称为 “基准”（pivot）;
//
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//
//递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
//
//递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。虽然一直递归下去，但是这个算法总会退出，因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

// left,right分别为数组的最左、最右下标
func quickSort(arr []int, left, right int) {
	if left >= right { // 递归截止条件
		return
	}
	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

// 分区函数
func partition(arr []int, left, right int) int {
	pivot := arr[right] // 选择数组的最后一个元素作为基准值
	i := left           // 数组最左元素的下标
	for j := left; j < right; j++ {
		if arr[j] < pivot { // 将小于基准值的元素交换到左边
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// 处理基准值
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// 空间复杂度：O(1), 在分区时使用原地交换的方式
