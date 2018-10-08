package heap

import "github.com/xiaomeng79/go-algorithm/data-structures/heap"

//堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法
//堆排序的平均时间复杂度为 Ο(nlogn)

func HeapSort(arr []int) []int {
	h := heap.NewMin()
	for i := 0; i < len(arr); i++ {
		h.Insert(heap.Int(arr[i]))
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = int(h.Extract().(heap.Int))
	}
	return arr
}
