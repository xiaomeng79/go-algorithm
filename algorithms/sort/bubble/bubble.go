//冒泡排序
package bubble

//原理：重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成

//算法步骤:
//1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
//2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
//3. 针对所有的元素重复以上的步骤，除了最后一个。
//4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

//时间复杂度:O(n²)

//推荐
func BubbleSort(arr []int) {
	for itemCount := len(arr) - 1; ; itemCount-- {
		swap := false //下一次判断是否发生交换，没交换，说明已经有序
		for i := 1; i <= itemCount; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}

//func BubbleSort(arr []int) []int {
//	length := len(arr)            //获取最大长度
//	for i := 0; i < length; i++ { //比较次数
//		for j := 0; j < length-1-i; j++ { //比较长度
//			if arr[j] > arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//	return arr
//}

//func BubbleSort2(arr []int) {
//	l := len(arr)
//	for i := 0; i<l; i++ {
//		flag := true
//		for j:=0;j<l-i-1;j++ {
//			if arr[j+1] < arr[j] {
//				arr[j+1],arr[j] = arr[j],arr[j+1]
//				flag = false
//			}
//		}
//		if flag {
//			break
//		}
//	}
//}