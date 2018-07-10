//插入排序
package insertion

//原理：通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
//稳定性：稳定
//算法步骤:
// 1.第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列
// 2.从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面

//[]int{3}
//[]int{3,44,56,38,77,38,26}

func InsertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}
