package testdata

var (
	Values = []struct {
		Nosort []int
		Sort   []int
	}{
		{[]int{3, 44, 56, 38, 77, 38, 26}, []int{3, 26, 38, 38, 44, 56, 77}},
		{[]int{3}, []int{3}},
		{[]int{3, -1, -6, 34, -78}, []int{-78, -6, -1, 3, 34}},
	}
)
