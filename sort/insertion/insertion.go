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

func InsertionSort2(arr []int) []int {
	length := len(arr)
	//if length == 1 {return arr}

	for i := 1; i < length; i++ {
		backup := arr[i]
		j := i -1;
		//将选出的被排数比较后插入左边有序区
		for  j >= 0 && backup < arr[j] {//注意j >= 0必须在前边，否则会数组越界
			arr[j+1] = arr[j]//移动有序数组
			j -- //反向移动下标
		}
		arr[j + 1] = backup //插队插入移动后的空位
	}
	return arr
}
