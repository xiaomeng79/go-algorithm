//选择排序
package selection

//原理：每一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置
//稳定: 不稳定
//适用: 数据规模越小越好
//算法步骤:
//1. 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
//2. 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
//3. 重复第二步，直到所有元素均排序完毕

func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ { //注意这里使用length-1少一次，
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return arr
}
