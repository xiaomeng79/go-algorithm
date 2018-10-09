package merge

//归并排序:将两个的有序数列合并成一个有序数列

//算法步骤（从上到下）:
//1.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
//2.设定两个指针，最初位置分别为两个已经排序序列的起始位置；
//3.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
//4.重复步骤 3 直到某一指针达到序列尾；
//5.将另一序列剩下的所有元素直接复制到合并序列尾。

//第一种方法
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	return merge(MergeSort(arr[0:mid]), MergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}

//第二种方法：推荐
func sort(arr []int) {
	var s = make([]int, len(arr)/2+1)
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2

	sort(arr[:mid])
	sort(arr[mid:])

	if arr[mid-1] <= arr[mid] {
		return
	}

	copy(s, arr[:mid])

	l, r := 0, mid

	for i := 0; ; i++ {
		if s[l] <= arr[r] {
			arr[i] = s[l]
			l++

			if l == mid {
				break
			}
		} else {
			arr[i] = arr[r]
			r++
			if r == len(arr) {
				copy(arr[i+1:], s[l:mid])
				break
			}
		}
	}
	return
}
