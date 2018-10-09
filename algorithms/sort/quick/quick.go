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
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quickSort(arr, left, partitionIndex-1)
		quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] { //这里是关键，找到一个比基准大的数和一个比基准小的树数进行交换
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1) //将基准交换到中间位置
	return index - 1

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
